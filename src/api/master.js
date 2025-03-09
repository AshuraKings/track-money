import { refreshToken } from './withauth'

export const getUsers = async () => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    const res = await fetch('/api/authed/users', { method: 'GET', headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await getUsers()
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

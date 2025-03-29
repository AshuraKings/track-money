export const logout = async () => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/logout')
    const res = await fetch('/api', { method: 'GET', headers })
    const status = res.status
    let resHeader = res.headers
    const body2 = await res.json()
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const authed = async () => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed')
    const res = await fetch('/api', { method: 'GET', headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await authed()
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const refreshToken = async () => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('refreshToken')}`)
    headers.set('ai-path', '/authed/refresh')
    const res = await fetch('/api', { method: 'GET', headers })
    const status = res.status
    let resHeader = res.headers
    const body2 = await res.json()
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}
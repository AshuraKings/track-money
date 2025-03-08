export const logout = async () => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    const res = await fetch('/api/logout', { method: 'GET', headers })
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
    const res = await fetch('/api/authed', { method: 'GET', headers })
    const status = res.status
    let resHeader = res.headers
    const body2 = await res.json()
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const refreshToken = async () => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('refreshToken')}`)
    const res = await fetch('/api/authed', { method: 'GET', headers })
    const status = res.status
    let resHeader = res.headers
    const body2 = await res.json()
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}
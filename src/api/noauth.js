export const register = async (body = { nm: '', username: '', password: '' }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    const res = await fetch('/api/register', { method: 'POST', body: JSON.stringify(body), headers })
    const status = res.status
    let resHeader = res.headers
    const body2 = await res.json()
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const login = async (body = { username: '', password: '' }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    const res = await fetch('/api/login', { method: 'POST', body: JSON.stringify(body), headers })
    const status = res.status
    let resHeader = res.headers
    const body2 = await res.json()
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}
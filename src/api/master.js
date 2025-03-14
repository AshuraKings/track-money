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

export const delUser = async (body = { id: 0 }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    const res = await fetch('/api/authed/users', { method: 'DELETE', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await delUser(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const delRole = async (body = { id: 0 }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    const res = await fetch('/api/authed/roles', { method: 'DELETE', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await delRole(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const editUser = async (body = { name: '', username: '', id: 0, role: 0 }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    const res = await fetch('/api/authed/users', { method: 'PUT', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await editUser(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const addUser = async (body = { name: '', username: '', password: '', role: 0 }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    const res = await fetch('/api/authed/users', { method: 'POST', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await addUser(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const editRole = async (body = { name: '', id: 0 }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    const res = await fetch('/api/authed/roles', { method: 'PUT', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await editRole(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const addRole = async (body = { name: '' }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    const res = await fetch('/api/authed/roles', { method: 'POST', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await addRole(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const getMenus = async () => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    const res = await fetch('/api/authed/menus', { method: 'GET', headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await getMenus()
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const getRoles = async () => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    const res = await fetch('/api/authed/roles', { method: 'GET', headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await getRoles()
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}
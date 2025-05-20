import { refreshToken } from './withauth'

export const getUsers = async () => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/users')
    const res = await fetch('/api', { method: 'GET', headers })
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
    headers.set('ai-path', '/authed/users')
    const res = await fetch('/api', { method: 'DELETE', body: JSON.stringify(body), headers })
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
    headers.set('ai-path', '/authed/roles')
    const res = await fetch('/api', { method: 'DELETE', body: JSON.stringify(body), headers })
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

export const delMenu = async (body = { id: 0 }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/menus')
    const res = await fetch('/api', { method: 'DELETE', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await delMenu(body)
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
    headers.set('ai-path', '/authed/users')
    const res = await fetch('/api', { method: 'PUT', body: JSON.stringify(body), headers })
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
    headers.set('ai-path', '/authed/users')
    const res = await fetch('/api', { method: 'POST', body: JSON.stringify(body), headers })
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
    headers.set('ai-path', '/authed/roles')
    const res = await fetch('/api', { method: 'PUT', body: JSON.stringify(body), headers })
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
    headers.set('ai-path', '/authed/roles')
    const res = await fetch('/api', { method: 'POST', body: JSON.stringify(body), headers })
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

export const editMenus = async (body = { label: '', createdAt: "2025-03-14T21:55:54.554924Z", id: 0 }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/menus')
    const res = await fetch('/api', { method: 'PUT', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await editMenus(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const addMenus = async (body = { label: '', createdAt: "2025-03-14T21:55:54.554924Z" }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/menus')
    const res = await fetch('/api', { method: 'POST', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await addMenus(body)
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
    headers.set('ai-path', '/authed/menus')
    const res = await fetch('/api', { method: 'GET', headers })
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
    headers.set('ai-path', '/authed/roles')
    const res = await fetch('/api', { method: 'GET', headers })
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

export const getRoleMenus = async (body = { id: 0 }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/roles/menus')
    const res = await fetch('/api?id=' + body.id, { method: 'GET', headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await getRoleMenus(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const addRoleMenus = async (body = { roleId: 0, menus: [] }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/roles/menus')
    const res = await fetch('/api', { method: 'POST', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await addRoleMenus(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const getWallets = async () => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/wallets')
    const res = await fetch('/api', { method: 'GET', headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await getWallets()
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const addWallet = async (body = { nm: '', balance: 0 }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/wallets')
    const res = await fetch('/api', { method: 'POST', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await addWallet(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const delWallet = async (body = { id: 0 }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/wallets')
    const res = await fetch('/api', { method: 'DELETE', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await delWallet(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const getIncomes = async () => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/incomes')
    const res = await fetch('/api', { method: 'GET', headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            const res = await getIncomes()
            return res
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const addIncome = async (body = { nm: '' }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/incomes')
    const res = await fetch('/api', { method: 'POST', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await addIncome(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const delIncome = async (body = { id: 0 }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/incomes')
    const res = await fetch('/api', { method: 'DELETE', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await delIncome(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const getExpenses = async () => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/expenses')
    const res = await fetch('/api', { method: 'GET', headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            const res = await getExpenses()
            return res
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const addExpenses = async (body = { nm: '' }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/expenses')
    const res = await fetch('/api', { method: 'POST', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await addExpenses(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const delExpenses = async (body = { id: 0 }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/expenses')
    const res = await fetch('/api', { method: 'DELETE', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await delExpenses(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const getFirstDateTransaksi = async () => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/monthly')
    const res = await fetch('/api', { method: 'GET', headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await getFirstDateTransaksi()
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const getChartOfMonth = async (month = '') => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/monthly')
    const res = await fetch('/api?month=' + month, { method: 'GET', headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await getChartOfMonth(month)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const getTransaksies = async (body = { page: 0, limit: 0 }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/transactions')
    const query = Object.keys(body).map(k => `${k}=${body[k]}`).join('&')
    const res = await fetch('/api?' + query, { method: 'GET', headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await getTransaksies(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

export const addTransaksies = async (body = { ket: '', date: '', amount: 0, admin: 0 }) => {
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    headers.set('Authorization', `Bearer ${localStorage.getItem('sessionToken')}`)
    headers.set('ai-path', '/authed/transactions')
    const res = await fetch('/api', { method: 'POST', body: JSON.stringify(body), headers })
    const status = res.status
    const body2 = await res.json()
    if (body2.msg === 'Token is expired') {
        const r = await refreshToken()
        const { headers, status } = r
        if (status >= 200 && status < 300) {
            localStorage.setItem('refreshToken', headers.refreshtoken)
            localStorage.setItem('sessionToken', headers.sessiontoken)
            return await addTransaksies(body)
        }
    }
    let resHeader = res.headers
    const h = {}
    resHeader.forEach((v, k) => h[k] = v)
    return { headers: h, body: body2, status }
}

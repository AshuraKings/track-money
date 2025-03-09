SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE IF NOT EXISTS roles(
    id BIGSERIAL PRIMARY KEY,
    nm VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP
);
MERGE INTO roles r USING (SELECT 'admin' nm) AS n ON r.nm=n.nm WHEN NOT MATCHED THEN INSERT(nm) VALUES(n.nm);
MERGE INTO roles r USING (SELECT 'fin' nm) AS n ON r.nm=n.nm WHEN NOT MATCHED THEN INSERT(nm) VALUES(n.nm);
MERGE INTO roles r USING (SELECT 'viewer' nm) AS n ON r.nm=n.nm WHEN NOT MATCHED THEN INSERT(nm) VALUES(n.nm);
MERGE INTO roles r USING (SELECT 'viewer-out' nm) AS n ON r.nm=n.nm WHEN NOT MATCHED THEN INSERT(nm) VALUES(n.nm);
CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    sandi VARCHAR(500) NOT NULL,
    nm VARCHAR(255) NOT NULL,
    role_id BIGINT NOT NULL REFERENCES roles(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS menus(
    id BIGSERIAL PRIMARY KEY,
    label VARCHAR(255) NOT NULL,
    link VARCHAR(500),
    icon VARCHAR(200),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP
);
MERGE INTO menus m USING (SELECT 'Dashboard' label,'/dashboard' link,'md-dashboard' icon) AS n ON n.label=m.label WHEN NOT MATCHED THEN INSERT(label,link,icon) VALUES(n.label,n.link,n.icon);
MERGE INTO menus m USING (SELECT 'Master' label,NULL link,'fa-database' icon) AS n ON n.label=m.label WHEN NOT MATCHED THEN INSERT(label,link,icon) VALUES(n.label,n.link,n.icon);
MERGE INTO menus m USING (SELECT 'Users' label,'/master/users' link,null icon) AS n ON n.label=m.label WHEN NOT MATCHED THEN INSERT(label,link,icon) VALUES(n.label,n.link,n.icon);
MERGE INTO menus m USING (SELECT 'Roles' label,'/master/roles' link,null icon) AS n ON n.label=m.label WHEN NOT MATCHED THEN INSERT(label,link,icon) VALUES(n.label,n.link,n.icon);
MERGE INTO menus m USING (SELECT 'Menus' label,'/master/menus' link,null icon) AS n ON n.label=m.label WHEN NOT MATCHED THEN INSERT(label,link,icon) VALUES(n.label,n.link,n.icon);
MERGE INTO menus m USING (SELECT 'Wallets' label,'/master/wallets' link,null icon) AS n ON n.label=m.label WHEN NOT MATCHED THEN INSERT(label,link,icon) VALUES(n.label,n.link,n.icon);
MERGE INTO menus m USING (SELECT 'Incomes' label,'/master/incomes' link,null icon) AS n ON n.label=m.label WHEN NOT MATCHED THEN INSERT(label,link,icon) VALUES(n.label,n.link,n.icon);
MERGE INTO menus m USING (SELECT 'Expenses' label,'/master/expenses' link,null icon) AS n ON n.label=m.label WHEN NOT MATCHED THEN INSERT(label,link,icon) VALUES(n.label,n.link,n.icon);
MERGE INTO menus m USING (SELECT 'Transactions' label,'/master/transactions' link,null icon) AS n ON n.label=m.label WHEN NOT MATCHED THEN INSERT(label,link,icon) VALUES(n.label,n.link,n.icon);
CREATE TABLE IF NOT EXISTS menu_has_menu(
    menu_id BIGINT NOT NULL REFERENCES menus(id),
    parent_id BIGINT NOT NULL REFERENCES menus(id)
);
MERGE INTO menu_has_menu m USING (SELECT 3 menu_id,2 parent_id) AS n ON n.menu_id=m.menu_id AND n.parent_id=m.parent_id WHEN NOT MATCHED THEN INSERT(menu_id,parent_id) VALUES(n.menu_id,n.parent_id);
MERGE INTO menu_has_menu m USING (SELECT 4 menu_id,2 parent_id) AS n ON n.menu_id=m.menu_id AND n.parent_id=m.parent_id WHEN NOT MATCHED THEN INSERT(menu_id,parent_id) VALUES(n.menu_id,n.parent_id);
MERGE INTO menu_has_menu m USING (SELECT 5 menu_id,2 parent_id) AS n ON n.menu_id=m.menu_id AND n.parent_id=m.parent_id WHEN NOT MATCHED THEN INSERT(menu_id,parent_id) VALUES(n.menu_id,n.parent_id);
MERGE INTO menu_has_menu m USING (SELECT 6 menu_id,2 parent_id) AS n ON n.menu_id=m.menu_id AND n.parent_id=m.parent_id WHEN NOT MATCHED THEN INSERT(menu_id,parent_id) VALUES(n.menu_id,n.parent_id);
MERGE INTO menu_has_menu m USING (SELECT 7 menu_id,2 parent_id) AS n ON n.menu_id=m.menu_id AND n.parent_id=m.parent_id WHEN NOT MATCHED THEN INSERT(menu_id,parent_id) VALUES(n.menu_id,n.parent_id);
MERGE INTO menu_has_menu m USING (SELECT 8 menu_id,2 parent_id) AS n ON n.menu_id=m.menu_id AND n.parent_id=m.parent_id WHEN NOT MATCHED THEN INSERT(menu_id,parent_id) VALUES(n.menu_id,n.parent_id);
MERGE INTO menu_has_menu m USING (SELECT 9 menu_id,2 parent_id) AS n ON n.menu_id=m.menu_id AND n.parent_id=m.parent_id WHEN NOT MATCHED THEN INSERT(menu_id,parent_id) VALUES(n.menu_id,n.parent_id);
CREATE TABLE IF NOT EXISTS role_has_menu(
    role_id BIGINT NOT NULL REFERENCES roles(id),
    menu_id BIGINT NOT NULL REFERENCES menus(id)
);
MERGE INTO role_has_menu r USING (SELECT 1 role_id,1 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
MERGE INTO role_has_menu r USING (SELECT 1 role_id,2 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
MERGE INTO role_has_menu r USING (SELECT 1 role_id,3 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
MERGE INTO role_has_menu r USING (SELECT 1 role_id,4 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
MERGE INTO role_has_menu r USING (SELECT 1 role_id,5 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
MERGE INTO role_has_menu r USING (SELECT 1 role_id,6 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
MERGE INTO role_has_menu r USING (SELECT 1 role_id,7 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
MERGE INTO role_has_menu r USING (SELECT 1 role_id,8 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
MERGE INTO role_has_menu r USING (SELECT 1 role_id,9 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
MERGE INTO role_has_menu r USING (SELECT 2 role_id,1 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
MERGE INTO role_has_menu r USING (SELECT 2 role_id,2 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
MERGE INTO role_has_menu r USING (SELECT 2 role_id,5 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
MERGE INTO role_has_menu r USING (SELECT 2 role_id,6 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
MERGE INTO role_has_menu r USING (SELECT 2 role_id,7 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
MERGE INTO role_has_menu r USING (SELECT 2 role_id,8 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
MERGE INTO role_has_menu r USING (SELECT 3 role_id,1 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
MERGE INTO role_has_menu r USING (SELECT 4 role_id,1 menu_id) AS n ON n.role_id=r.role_id AND n.menu_id=r.menu_id WHEN NOT MATCHED THEN INSERT(role_id,menu_id) VALUES(n.role_id,n.menu_id);
CREATE TABLE IF NOT EXISTS wallets(
    id BIGSERIAL PRIMARY KEY,
    nm VARCHAR(255) NOT NULL UNIQUE,
    balance DECIMAL NOT NULL DEFAULT 0.0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP
);
MERGE INTO wallets w USING (SELECT 'Saving' nm) AS n ON n.nm=w.nm WHEN NOT MATCHED THEN INSERT(nm) VALUES(n.nm);
MERGE INTO wallets w USING (SELECT 'Uang Saku Mas' nm) AS n ON n.nm=w.nm WHEN NOT MATCHED THEN INSERT(nm) VALUES(n.nm);
MERGE INTO wallets w USING (SELECT 'Dompet' nm) AS n ON n.nm=w.nm WHEN NOT MATCHED THEN INSERT(nm) VALUES(n.nm);
CREATE TABLE IF NOT EXISTS incomes(
    id BIGSERIAL PRIMARY KEY,
    nm VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP
);
MERGE INTO incomes i USING (SELECT 'Gaji Mas' nm) AS n ON n.nm=i.nm WHEN NOT MATCHED THEN INSERT(nm) VALUES(n.nm);
CREATE TABLE IF NOT EXISTS expenses(
    id BIGSERIAL PRIMARY KEY,
    nm VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS transaksi(
    kode VARCHAR(20) PRIMARY KEY,
    ket VARCHAR(500) NOT NULL,
    amount DECIMAL NOT NULL,
    admin_fee DECIMAL NOT NULL DEFAULT 0.0,
    trx_date DATE NOT NULL,
    from_wallet_id BIGINT REFERENCES wallets(id),
    to_wallet_id BIGINT REFERENCES wallets(id),
    income_id BIGINT REFERENCES incomes(id),
    expenses_id BIGINT REFERENCES expenses(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP
);
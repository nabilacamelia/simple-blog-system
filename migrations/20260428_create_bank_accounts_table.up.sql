CREATE TABLE bank_accounts (
    id SERIAL PRIMARY KEY,
    company_id INTEGER NOT NULL,
    bank_name VARCHAR(100) NOT NULL, -- Contoh: BCA, Mandiri
    account_number VARCHAR(50) NOT NULL, -- Nomor Rekening
    account_holder VARCHAR(255), -- Nama Pemilik Rekening
    branch_office VARCHAR(255), -- Cabang Bank
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_company FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE
);
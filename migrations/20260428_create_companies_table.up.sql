CREATE TABLE companies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL, -- Nama PT (PT. Zerra Teknologi Integrasi) [cite: 1, 68]
    legal_name VARCHAR(255),    -- Nama Legal Perusahaan
    address TEXT,               -- Alamat (Jl. Joyogrand) [cite: 2]
    phone VARCHAR(20),          -- Nomor Telepon [cite: 4]
    email VARCHAR(100),          -- Email [cite: 4]
    logo_url TEXT,              -- Link Logo Perusahaan
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
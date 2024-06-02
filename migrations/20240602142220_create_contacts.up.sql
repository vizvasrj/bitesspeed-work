CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(20),
    email VARCHAR(255),
    linked_id INT,
    link_precedence VARCHAR(10) DEFAULT 'primary',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);
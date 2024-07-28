CREATE TABLE IF NOT EXISTS sites (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    description TEXT NOT NULL UNIQUE,
    url TEXT NOT NULL UNIQUE
);

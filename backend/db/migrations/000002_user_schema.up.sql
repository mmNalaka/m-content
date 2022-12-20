CREATE TABLE users(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    name TEXT NOT NULL,
    avatar TEXT, -- NULL if no avatar
    location TEXT, -- NULL if no location
    title TEXT, -- NULL if no title
    bio TEXT, -- NULL if no bio
    status TEXT NOT NULL DEFAULT 'active', -- 'active', 'inactive', 'banned'
    tags JSON, -- NULL if no tags
    meta JSON, -- NULL if no meta

    last_login TIMESTAMPTZ, -- NULL if never logged in
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );
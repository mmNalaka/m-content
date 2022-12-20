CREATE TABLE collections(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    slug TEXT NOT NULL UNIQUE,
    name TEXT,
    notes TEXT,
    singleton BOOLEAN DEFAULT FALSE NOT NULL,
    schema JSON NOT NULL,

    listRule TEXT,
    viewRule TEXT,
    createRule TEXT,
    updateRule TEXT,
    deleteRule TEXT,

    tags JSON, -- NULL if no tags
    meta JSON, -- NULL if no meta

    created_by UUID,
    updated_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

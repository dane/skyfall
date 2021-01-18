CREATE TABLE accounts (
	pk BIGSERIAL PRIMARY KEY,
	id UUID NOT NULL DEFAULT UUID_GENERATE_V4(),
	name TEXT NOT NULL,
	password_hash TEXT NOT NULL DEFAULT '',
	password_salt TEXT NOT NULL DEFAULT '',
	properties JSONB DEFAULT '{}',
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	verified_at TIMESTAMPTZ,
	suspended_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

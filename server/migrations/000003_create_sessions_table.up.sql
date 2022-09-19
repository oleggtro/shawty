CREATE TABLE IF NOT EXISTS sessions (
    token UUID NOT NULL DEFAULT gen_random_uuid(),
    subject UUID UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp + (5 * interval '1 minute')
);
CREATE TABLE IF NOT EXISTS redirects (
    id TEXT PRIMARY KEY,
    redirect_to TEXT NOT NULL,
    owner UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
    uses INTEGER NOT NULL DEFAULT 0
);
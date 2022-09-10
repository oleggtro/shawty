CREATE TABLE IF NOT EXISTS redirects(
    id TEXT PRIMARY KEY,
    redirect_to TEXT NOT NULL,
    owner UUID NOT NULL
);
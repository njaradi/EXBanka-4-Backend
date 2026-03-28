CREATE TABLE activation_tokens (
    token       VARCHAR PRIMARY KEY,
    employee_id BIGINT      NOT NULL,
    expires_at  TIMESTAMPTZ NOT NULL
);

-- Seed tokens for Cypress e2e tests (celina1 S8/S9)
INSERT INTO activation_tokens (token, employee_id, expires_at) VALUES
  ('cypress-valid-token-s8',   100, '2099-12-31 23:59:59+00'),
  ('cypress-expired-token-s9', 101, '2020-01-01 00:00:00+00')
ON CONFLICT (token) DO NOTHING;

CREATE TABLE IF NOT EXISTS password_reset_tokens (
    token       VARCHAR PRIMARY KEY,
    employee_id BIGINT      NOT NULL,
    expires_at  TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS client_activation_tokens (
    token      VARCHAR PRIMARY KEY,
    client_id  BIGINT      NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS two_factor_approvals (
    id          BIGSERIAL    PRIMARY KEY,
    client_id   BIGINT       NOT NULL,
    action_type VARCHAR      NOT NULL,
    payload     TEXT         NOT NULL DEFAULT '{}',
    status      VARCHAR      NOT NULL DEFAULT 'PENDING',
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT now(),
    expires_at  TIMESTAMPTZ  NOT NULL DEFAULT now() + interval '5 minutes'
);

CREATE TABLE IF NOT EXISTS push_tokens (
    client_id   BIGINT   PRIMARY KEY,
    token       VARCHAR  NOT NULL
);

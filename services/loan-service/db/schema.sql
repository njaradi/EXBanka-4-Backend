CREATE TABLE IF NOT EXISTS loans (
    id                      BIGSERIAL PRIMARY KEY,
    loan_number             BIGINT      NOT NULL UNIQUE,
    account_number          VARCHAR     NOT NULL,
    client_id               BIGINT      NOT NULL,
    loan_type               VARCHAR     NOT NULL, -- gotovinski, stambeni, auto, refinansirajuci, studentski
    interest_rate_type      VARCHAR     NOT NULL, -- fiksna, varijabilna
    amount                  NUMERIC(20, 2) NOT NULL,
    currency                VARCHAR     NOT NULL,
    repayment_period        INT         NOT NULL, -- total number of installments
    nominal_rate            NUMERIC(10, 4) NOT NULL, -- base rate at creation (annual %)
    effective_rate          NUMERIC(10, 4) NOT NULL, -- current effective rate (annual %)
    agreed_date             DATE        NOT NULL DEFAULT CURRENT_DATE,
    maturity_date           DATE        NOT NULL,
    next_installment_amount NUMERIC(20, 2),
    next_installment_date   DATE,
    remaining_debt          NUMERIC(20, 2),
    status                  VARCHAR     NOT NULL DEFAULT 'PENDING', -- PENDING, APPROVED, REJECTED, PAID_OFF, IN_DELAY
    purpose                 VARCHAR,
    monthly_salary          NUMERIC(20, 2),
    employment_status       VARCHAR,    -- stalno, privremeno, nezaposlen
    employment_period       INT,        -- months at current employer
    contact_phone           VARCHAR
);

CREATE TABLE IF NOT EXISTS loan_installments (
    id                  BIGSERIAL PRIMARY KEY,
    loan_id             BIGINT      NOT NULL REFERENCES loans(id),
    installment_amount  NUMERIC(20, 2) NOT NULL,
    interest_rate       NUMERIC(10, 4) NOT NULL, -- rate at time of payment
    currency            VARCHAR     NOT NULL,
    expected_due_date   DATE        NOT NULL,
    actual_due_date     DATE,
    status              VARCHAR     NOT NULL DEFAULT 'UNPAID', -- PAID, UNPAID, LATE
    retry_count         INT         NOT NULL DEFAULT 0,
    last_retry_at       TIMESTAMPTZ
);

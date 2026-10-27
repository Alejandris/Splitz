ALTER TABLE public.users
    ADD COLUMN IF NOT EXISTS currency_code VARCHAR(3) NOT NULL DEFAULT 'MXN';

ALTER TABLE public.periods
    ADD COLUMN IF NOT EXISTS start_date DATE,
    ADD COLUMN IF NOT EXISTS end_date DATE,
    ADD COLUMN IF NOT EXISTS currency_code VARCHAR(3) NOT NULL DEFAULT 'MXN';

UPDATE public.periods
SET start_date = make_date(year, month, 1),
    end_date = (make_date(year, month, 1) + INTERVAL '1 month - 1 day')::date
WHERE start_date IS NULL OR end_date IS NULL;

ALTER TABLE public.periods
    ALTER COLUMN start_date SET NOT NULL,
    ALTER COLUMN end_date SET NOT NULL,
    ADD CONSTRAINT periods_date_range CHECK (end_date >= start_date),
    ADD CONSTRAINT periods_currency_code CHECK (currency_code IN ('MXN', 'COP', 'USD', 'EUR'));

ALTER TABLE public.users
    ADD CONSTRAINT users_currency_code CHECK (currency_code IN ('MXN', 'COP', 'USD', 'EUR'));

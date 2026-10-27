ALTER TABLE public.categories
    ADD COLUMN IF NOT EXISTS created_by_user_id UUID REFERENCES public.users(id) ON DELETE CASCADE,
    ADD COLUMN IF NOT EXISTS is_active BOOLEAN NOT NULL DEFAULT TRUE;

CREATE INDEX IF NOT EXISTS idx_categories_owner_active
    ON public.categories (created_by_user_id, is_active);

INSERT INTO public.categories (name, type, created_by_user_id, is_active)
SELECT seed.name, seed.type, NULL, TRUE
FROM (VALUES
    ('Necesidades', 'EXPENSE'),
    ('Vivienda', 'EXPENSE'),
    ('Servicios', 'EXPENSE'),
    ('Transporte', 'EXPENSE'),
    ('Alimentación', 'EXPENSE'),
    ('Deudas', 'DEBT'),
    ('Ahorro', 'SAVING'),
    ('Inversión', 'SAVING'),
    ('Deseos', 'EXPENSE'),
    ('Entretenimiento', 'EXPENSE'),
    ('Salud', 'EXPENSE'),
    ('Educación', 'EXPENSE')
) AS seed(name, type)
WHERE NOT EXISTS (
    SELECT 1 FROM public.categories c
    WHERE c.name = seed.name AND c.created_by_user_id IS NULL
);

INSERT INTO public.methods (name, is_custom, created_by_user_id)
SELECT seed.name, FALSE, NULL
FROM (VALUES
    ('Tradicional 50/30/20'),
    ('Modo Enfoque'),
    ('Modo Inversionista'),
    ('Modo Disfrute')
) AS seed(name)
WHERE NOT EXISTS (
    SELECT 1 FROM public.methods m
    WHERE m.name = seed.name AND m.created_by_user_id IS NULL
);

INSERT INTO public.method_allocations (method_id, category_id, percentage)
SELECT m.id, c.id, seed.percentage
FROM (VALUES
    ('Tradicional 50/30/20', 'Necesidades', 50.00),
    ('Tradicional 50/30/20', 'Deseos', 30.00),
    ('Tradicional 50/30/20', 'Ahorro', 20.00),
    ('Modo Enfoque', 'Necesidades', 50.00),
    ('Modo Enfoque', 'Deudas', 30.00),
    ('Modo Enfoque', 'Ahorro', 10.00),
    ('Modo Enfoque', 'Deseos', 10.00),
    ('Modo Inversionista', 'Necesidades', 40.00),
    ('Modo Inversionista', 'Ahorro', 40.00),
    ('Modo Inversionista', 'Deseos', 20.00),
    ('Modo Disfrute', 'Necesidades', 40.00),
    ('Modo Disfrute', 'Deseos', 40.00),
    ('Modo Disfrute', 'Ahorro', 20.00)
) AS seed(method_name, category_name, percentage)
JOIN public.methods m ON m.name = seed.method_name AND m.created_by_user_id IS NULL
JOIN public.categories c ON c.name = seed.category_name AND c.created_by_user_id IS NULL
WHERE NOT EXISTS (
    SELECT 1 FROM public.method_allocations ma
    WHERE ma.method_id = m.id AND ma.category_id = c.id
);

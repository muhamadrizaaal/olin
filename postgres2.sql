CREATE INDEX idx_orders_user_id ON second.public.orders(user_id);

SELECT u.name, o.amount, o.created_at
FROM first.public.users u
JOIN second.public.orders o ON u.id = o.user_id;

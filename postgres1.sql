CREATE INDEX idx_orders_user_id ON orders(user_id);
CREATE INDEX idx_orders_created_at ON orders(created_at);

SELECT u.name, SUM(o.amount) AS total_spent
FROM users u
JOIN orders o ON u.id = o.user_id
WHERE o.created_at >= '2022-01-01'
GROUP BY u.id, u.name
HAVING SUM(o.amount) >= 1000;

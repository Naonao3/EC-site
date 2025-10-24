-- ==========================================
-- シードデータ投入SQL
-- ==========================================

-- ==========================================
-- 1. カテゴリデータ（柴犬画像用）
-- ==========================================

-- 親カテゴリ
INSERT INTO categories (name, slug, parent_id, created_at) VALUES
('表情', 'expression', NULL, CURRENT_TIMESTAMP),
('シチュエーション', 'situation', NULL, CURRENT_TIMESTAMP),
('季節', 'season', NULL, CURRENT_TIMESTAMP),
('ポーズ', 'pose', NULL, CURRENT_TIMESTAMP);

-- 子カテゴリ（表情）
INSERT INTO categories (name, slug, parent_id, created_at) VALUES
('笑顔', 'smile', (SELECT id FROM categories WHERE slug = 'expression'), CURRENT_TIMESTAMP),
('真剣', 'serious', (SELECT id FROM categories WHERE slug = 'expression'), CURRENT_TIMESTAMP),
('眠そう', 'sleepy', (SELECT id FROM categories WHERE slug = 'expression'), CURRENT_TIMESTAMP),
('驚き', 'surprised', (SELECT id FROM categories WHERE slug = 'expression'), CURRENT_TIMESTAMP);

-- 子カテゴリ（シチュエーション）
INSERT INTO categories (name, slug, parent_id, created_at) VALUES
('散歩', 'walk', (SELECT id FROM categories WHERE slug = 'situation'), CURRENT_TIMESTAMP),
('遊び', 'play', (SELECT id FROM categories WHERE slug = 'situation'), CURRENT_TIMESTAMP),
('お昼寝', 'nap', (SELECT id FROM categories WHERE slug = 'situation'), CURRENT_TIMESTAMP),
('食事', 'meal', (SELECT id FROM categories WHERE slug = 'situation'), CURRENT_TIMESTAMP);

-- 子カテゴリ（季節）
INSERT INTO categories (name, slug, parent_id, created_at) VALUES
('春', 'spring', (SELECT id FROM categories WHERE slug = 'season'), CURRENT_TIMESTAMP),
('夏', 'summer', (SELECT id FROM categories WHERE slug = 'season'), CURRENT_TIMESTAMP),
('秋', 'autumn', (SELECT id FROM categories WHERE slug = 'season'), CURRENT_TIMESTAMP),
('冬', 'winter', (SELECT id FROM categories WHERE slug = 'season'), CURRENT_TIMESTAMP);

-- 子カテゴリ（ポーズ）
INSERT INTO categories (name, slug, parent_id, created_at) VALUES
('座る', 'sitting', (SELECT id FROM categories WHERE slug = 'pose'), CURRENT_TIMESTAMP),
('立つ', 'standing', (SELECT id FROM categories WHERE slug = 'pose'), CURRENT_TIMESTAMP),
('走る', 'running', (SELECT id FROM categories WHERE slug = 'pose'), CURRENT_TIMESTAMP),
('寝る', 'lying', (SELECT id FROM categories WHERE slug = 'pose'), CURRENT_TIMESTAMP);

-- ==========================================
-- 2. テストユーザー
-- ==========================================

-- パスワードは全て "password123" のハッシュ値
-- 実際の運用では bcrypt でハッシュ化したものを使用
INSERT INTO users (email, password_hash, name, created_at, updated_at) VALUES
('demo@example.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'デモユーザー', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('admin@example.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '管理者', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('test@example.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'テストユーザー', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- ==========================================
-- 3. サンプル商品データ（後日実際の柴犬画像で置き換え）
-- ==========================================

-- サンプル商品（表情カテゴリ）
INSERT INTO products (name, description, price, category_id, image_url, created_at, updated_at) VALUES
('笑顔の柴犬 - 春の散歩', '満開の桜の下で笑顔の柴犬。春の陽気な雰囲気が伝わる1枚。', 800,
 (SELECT id FROM categories WHERE slug = 'smile'),
 'https://via.placeholder.com/800x600?text=Shiba+Smile',
 CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),

('真剣な表情の柴犬', '何かを見つめる真剣な眼差し。凛々しい柴犬の魅力。', 900,
 (SELECT id FROM categories WHERE slug = 'serious'),
 'https://via.placeholder.com/800x600?text=Shiba+Serious',
 CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),

('眠そうな柴犬', 'うとうとしている可愛い瞬間を捉えた癒しの1枚。', 750,
 (SELECT id FROM categories WHERE slug = 'sleepy'),
 'https://via.placeholder.com/800x600?text=Shiba+Sleepy',
 CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- サンプル商品（シチュエーションカテゴリ）
INSERT INTO products (name, description, price, category_id, image_url, created_at, updated_at) VALUES
('散歩中の柴犬', '公園を楽しそうに歩く柴犬。元気いっぱいの様子。', 850,
 (SELECT id FROM categories WHERE slug = 'walk'),
 'https://via.placeholder.com/800x600?text=Shiba+Walk',
 CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),

('ボール遊び', 'ボールを追いかける躍動感ある瞬間。', 900,
 (SELECT id FROM categories WHERE slug = 'play'),
 'https://via.placeholder.com/800x600?text=Shiba+Play',
 CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- サンプル商品（季節カテゴリ）
INSERT INTO products (name, description, price, category_id, image_url, created_at, updated_at) VALUES
('桜と柴犬', '満開の桜の下で佇む柴犬。春らしい穏やかな1枚。', 1000,
 (SELECT id FROM categories WHERE slug = 'spring'),
 'https://via.placeholder.com/800x600?text=Shiba+Spring',
 CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),

('夏の海辺', '砂浜で楽しそうにする柴犬。夏の思い出。', 950,
 (SELECT id FROM categories WHERE slug = 'summer'),
 'https://via.placeholder.com/800x600?text=Shiba+Summer',
 CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),

('紅葉と柴犬', '色づいた紅葉を背景にした秋らしい1枚。', 1000,
 (SELECT id FROM categories WHERE slug = 'autumn'),
 'https://via.placeholder.com/800x600?text=Shiba+Autumn',
 CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),

('雪遊び', '雪の中ではしゃぐ柴犬。冬ならではの可愛さ。', 1100,
 (SELECT id FROM categories WHERE slug = 'winter'),
 'https://via.placeholder.com/800x600?text=Shiba+Winter',
 CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- ==========================================
-- 4. 在庫データ（各商品に対応）
-- ==========================================

INSERT INTO inventory (product_id, stock_quantity, reserved_quantity, updated_at)
SELECT id, 10, 0, CURRENT_TIMESTAMP
FROM products;

-- ==========================================
-- 確認用クエリ
-- ==========================================

-- カテゴリ確認
SELECT
    c1.id,
    c1.name AS category_name,
    c1.slug,
    c2.name AS parent_name
FROM categories c1
LEFT JOIN categories c2 ON c1.parent_id = c2.id
ORDER BY c1.parent_id NULLS FIRST, c1.id;

-- ユーザー確認
SELECT id, email, name, created_at FROM users;

-- 商品確認
SELECT
    p.id,
    p.name,
    p.price,
    c.name AS category_name,
    i.stock_quantity
FROM products p
JOIN categories c ON p.category_id = c.id
JOIN inventory i ON p.id = i.product_id
ORDER BY p.id;
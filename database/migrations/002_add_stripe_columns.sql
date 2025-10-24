-- ==========================================
-- Stripe統合用カラム追加マイグレーション
-- ==========================================

-- paymentsテーブルにStripe関連カラムを追加
ALTER TABLE payments
ADD COLUMN stripe_payment_intent_id VARCHAR(100),
ADD COLUMN stripe_payment_method_id VARCHAR(100);

-- インデックス追加（検索高速化）
CREATE INDEX idx_payments_stripe_payment_intent ON payments(stripe_payment_intent_id);
CREATE INDEX idx_payments_stripe_payment_method ON payments(stripe_payment_method_id);

-- コメント追加
COMMENT ON COLUMN payments.stripe_payment_intent_id IS 'Stripe Payment Intent ID';
COMMENT ON COLUMN payments.stripe_payment_method_id IS 'Stripe Payment Method ID';

-- マイグレーション完了確認用
SELECT
    table_name,
    column_name,
    data_type
FROM information_schema.columns
WHERE table_name = 'payments'
    AND column_name LIKE 'stripe%'
ORDER BY ordinal_position;
SELECT wallet.wallet_id as wallet_id,
       wallet.username as username , 
       wallet.wallet_description as wallet_description,
       wallet_record.record_name as wallet_record_name
FROM user 
JOIN wallet 
    ON wallet.user_id=user.id 
LEFT JOIN wallet_record 
    ON wallet_record.wallet_id=wallet.wallet_id 
WHERE user.id = ? 
ORDER BY CAST(wallet_record.wallet_id AS UNSIGNED)
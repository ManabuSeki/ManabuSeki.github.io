	req, err := http.NewRequestWithContext(ctx, http.MethodGet, originalURL.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Amz-Expires", expireStr)

	now := time.Now()

	if err := signer.SignHTTP(ctx, cred, req, "UNSIGNED-PAYLOAD", "s3", cfg.Region, now); err != nil {
		return err
	}

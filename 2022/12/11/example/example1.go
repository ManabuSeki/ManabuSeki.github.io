	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("ap-northeast-1"))
	if err != nil {
		return err
	}

	signer := v4.NewSigner(func(signer *v4.SignerOptions) {
		signer.DisableURIPathEscaping = true
	})
	cred, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		return err
	}

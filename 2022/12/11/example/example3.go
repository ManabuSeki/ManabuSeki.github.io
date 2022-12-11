headers, err := GetSignHeader(ctx, url, http.MethodGet, sign.UnsignedPayload, "s3") // main.goの処理をメソッド化
if err != nil {
	return nil, err
}
headers.Add("signed-url", url)            // S3のURL
headers.Add("X-Accel-Redirect", "/files") // 今回用意したnginxX-Accel-Redirect用のパス
headers.Add("X-Amz-Content-Sha256", "UNSIGNED-PAYLOAD")

for k := range headers {
	ctx.ResponseData.Header().Set(k, headers.Get(k))
}

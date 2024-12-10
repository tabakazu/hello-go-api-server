# OpenAPI の管理に関するドキュメント

## OpenApi Generator

作業は本リポジトリのルートディレクトリで行う

### 結合された仕様書の生成

```bash
docker run -u ${UID} --rm -v ${PWD}:/local \
	openapitools/openapi-generator-cli:v7.7.0 generate \
	--input-spec /local/api/openapi/specification/swagger.yaml \
	--generator-name openapi-yaml \
	--output /local/api/openapi/dist \
	--additional-properties outputFile=swagger.gen.yaml
```

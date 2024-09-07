The following source configuration file will sync to a PostgreSQL database. See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.


You can use Basic authentication or use Jamf's Client Credentials OAuth flow. See Jamf's documentation on how to get a Client ID and Client Secret. 

```yaml
kind: source
spec:
  name: "jamf"
  path: "jsifuentes/jamf"
  registry: "cloudquery"
  version: "v1.0.0"
  destinations:
    - "postgresql"
  spec:
    instance_domain: ${JAMF_INSTANCE_DOMAIN}
    auth_method: oauth # or "basic"
    # if oauth:
    # client_id: ${JAMF_CLIENT_ID}
    # client_secret: ${JAMF_CLIENT_SECRET}
    # if basic:
    # username: ${JAMF_USERNAME}
    # password: ${JAMF_PASSWORD}
```

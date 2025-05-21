Directory Surf
===============




## Edit `.env`

To run Directory Surf, create a `.env` file from the example:

```bash
cp .env.example .env
```

Then, fill in the required values as described below:


### App Base Info

* **`APP_NAME`** – Name of your app (displayed in UI/logs). You choose this.
* **`APP_PORT`** – Port your app runs on locally (e.g., `3000`). Common default: 3000.
* **`APP_DOMAIN`** – Your app’s full domain (e.g., `https://myapp.com`). Used in auth/cookies/emails.
* **`ADMIN_EMAIL`** – Email of the app admin; used for contact or admin access.


### Database

* **`DATABASE_URL`** – Full PostgreSQL connection string (e.g., `postgres://user:pass@host:port/db`).
  🔍 *Find this in your PostgreSQL provider or your local DB setup.*



### Stripe (Payments)

* **`STRIPE_SECRET_KEY`** – Your Stripe API secret key for backend operations.
  🔍 *Get it from [Stripe Dashboard > Developers > API Keys](https://dashboard.stripe.com/apikeys).*

* **`STRIPE_WEBHOOK_SECRET`** – Secret for verifying Stripe webhooks.
  🔍 *Found in [Stripe Dashboard > Developers > Webhooks](https://dashboard.stripe.com/webhooks) after setting up a webhook endpoint.*


### Cloudflare R2 (Object Storage)

* **`CLOUDFLARE_R2_ACCESS_KEY` / `SECRET_KEY`** – API credentials for accessing R2 storage.
  🔍 *Found in [Cloudflare Dashboard > R2 > API Tokens](https://dash.cloudflare.com).*

* **`CLOUDFLARE_R2_BUCKET_NAME`** – Name of your R2 bucket.
  🔍 *Set when creating the bucket in Cloudflare R2.*

* **`CLOUDFLARE_R2_ENDPOINT`** – Public endpoint for your R2 bucket.
  🔍 *Provided by Cloudflare when setting up R2.*


### Email Service

* **`RESEND_API_KEY`** – API key for sending emails with [Resend](https://resend.com).
  🔍 *Get it from your Resend dashboard.*


### CAPTCHA – Cloudflare Turnstile

* **`TURNSTILE_SITE_KEY`** – Public key for embedding Turnstile CAPTCHA.
  🔍 *Available from [Cloudflare Turnstile Dashboard](https://dash.cloudflare.com).*

* **`TURNSTILE_SECRET_KEY`** – Secret for backend CAPTCHA verification.
  🔍 *Also found in Turnstile settings.*


### 🔐 OAuth – GitHub

* **`GITHUB_CLIENT_ID` / `CLIENT_SECRET`** – Credentials for GitHub OAuth login.
  🔍 *Create at [GitHub Developer Settings > OAuth Apps](https://github.com/settings/developers).*


### 🔐 OAuth – Google

* **`GOOGLE_CLIENT_ID` / `CLIENT_SECRET`** – Credentials for Google OAuth login.
  🔍 *Create at [Google Cloud Console > Credentials](https://console.cloud.google.com/apis/credentials).*
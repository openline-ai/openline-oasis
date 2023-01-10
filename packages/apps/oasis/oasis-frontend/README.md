This is a [Next.js](https://nextjs.org/) project bootstrapped
with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app).

## Getting Started

### Setting up SSO

In fusion auth create an app for "oasis" then the following keys need to be updated in .env.local

* NEXTAUTH_OAUTH_CLIENT_ID
* NEXTAUTH_OAUTH_CLIENT_SECRET
* NEXTAUTH_OAUTH_TENANT_ID

### Setting up SSO (local k8s)

modify the env vars in the following file

```
deployment/k8s/local-minikube/apps-config/oasis-frontend.yaml
```

then modify your /etc/hosts as follows to resolve fusionauth-customer-os.openline.svc.cluster.local

```
127.0.0.1	localhost fusionauth-customer-os.openline.svc.cluster.local
```

###Install dependencies:

```bash

npm install

```

Run the development server:

```bash
npm run dev
# or
yarn dev
```

To run inside of docker:

```bash
docker buildx build -t ghcr.io/openline-ai/openline-oasis/oasis-frontend-dev --platform linux/amd64 --build-arg NODE_ENV=dev .
docker run -p 3006:3006 ghcr.io/openline-ai/openline-oasis/oasis-frontend-dev
```

Open [http://localhost:3006](http://localhost:3006) with your browser to see the result.

You can start editing the page by modifying `pages/index.tsx`. The page auto-updates as you edit the file.

[API routes](https://nextjs.org/docs/api-routes/introduction) can be accessed
on [http://localhost:3000/api/hello](http://localhost:3000/api/hello). This endpoint can be edited
in `pages/api/hello.ts`.

The `pages/api` directory is mapped to `/api/*`. Files in this directory are treated
as [API routes](https://nextjs.org/docs/api-routes/introduction) instead of React pages.

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js/) - your feedback and contributions
are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use
the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme)
from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/deployment) for more details.



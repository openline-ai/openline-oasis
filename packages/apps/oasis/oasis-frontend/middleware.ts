import {withAuth, NextRequestWithAuth} from "next-auth/middleware"
import {NextRequest, NextResponse} from "next/server";

// More on how NextAuth.js middleware works: https://next-auth.js.org/configuration/nextjs#middleware

export default withAuth(function middleware(request: NextRequestWithAuth) {
        var newURL = '';

        if (request.nextUrl.pathname.startsWith('/oasis-api/')) {
            newURL = process.env.NEXT_PUBLIC_OASIS_API_PATH + "/" + request.nextUrl.pathname.substring(("/oasis-api/").length);
        } else if(request.nextUrl.pathname.startsWith('/customer-os-api/')){
            newURL = process.env.NEXT_PUBLIC_CUSTOMER_OS_API_PATH + "/" + request.nextUrl.pathname.substring(("/customer-os-api/").length);
        }

        if (request.nextUrl.searchParams) {
            newURL = newURL + "?" + request.nextUrl.searchParams.toString()
        }
        console.log("Rewriting url to " + newURL);
        console.log("middleware: " + JSON.stringify(request.nextauth.token));

        const requestHeaders = new Headers(request.headers);

        if (request.nextUrl.pathname.startsWith('/oasis-api')) {
            requestHeaders.set('X-Openline-API-KEY', process.env.OASIS_API_KEY as string)
        } else if(request.nextUrl.pathname.startsWith('/customer-os-api')){
            requestHeaders.set('X-Openline-API-KEY', process.env.CUSTOMER_OS_API_KEY as string)
        }

        return NextResponse.rewrite(new URL(newURL, request.url),
            {
                request: {
                    headers: requestHeaders,
                },
            }
        )

    },
    {
        callbacks: {
            authorized({req, token}) {
                console.log("Got Token: " + JSON.stringify(token));
                if (token) return true
                return false;
            },
        },
    })

export const config = {
    matcher: ['/oasis-api/(.*)', '/customer-os-api/(.*)'],
}
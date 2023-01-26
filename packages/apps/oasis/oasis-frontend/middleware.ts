import {NextRequest, NextResponse} from "next/server";

export function middleware(request: NextRequest) {
    var newURL = '';

    if (request.nextUrl.pathname.startsWith('/oasis-api/')) {
        newURL = process.env.NEXT_PUBLIC_OASIS_API_PATH + "/" + request.nextUrl.pathname.substring(("/oasis-api/").length);
    } else if (request.nextUrl.pathname.startsWith('/customer-os-api/')) {
        newURL = process.env.NEXT_PUBLIC_CUSTOMER_OS_API_PATH + "/" + request.nextUrl.pathname.substring(("/customer-os-api/").length);
    } else {
        return NextResponse.next();
    }

    return fetch(`${process.env.ORY_SDK_URL}/sessions/whoami`, {
        headers: {
            cookie: request.headers.get("cookie") || "",
        },
    }).then((resp) => {
        // there must've been no response (invalid URL or something...)
        if (!resp) {
            console.log("no response");
            return NextResponse.redirect(new URL("/api/.ory/ui/login", request.url))
        }

        // the user is not signed in
        if (resp.status === 401) {
            console.log("not signed in");
            return NextResponse.redirect(new URL("/api/.ory/ui/login", request.url))
        }

        console.log("User is signed in, redirecting to " + newURL);

        return getRedirectUrl(newURL, request);
    }).catch((err) => {
        console.log(`Global Session Middleware error: ${JSON.stringify(err)}`)
        if (!err.response) {
            console.log("no response");
            return NextResponse.redirect(new URL("/api/.ory/ui/login", request.url))
        }
        switch (err.response?.status) {
            // 422 we need to redirect the user to the location specified in the response
            case 422:
                console.log("422");
                return NextResponse.redirect(new URL("/api/.ory/ui/login", request.url))
            //return router.push("/login", { query: { aal: "aal2" } })
            case 401:
                console.log("401");
                // The user is not logged in, so we redirect them to the login page.
                return NextResponse.redirect(new URL("/api/.ory/ui/login", request.url))
            case 404:
                console.log("404");
                // the SDK is not configured correctly
                // we set this up so you can debug the issue in the browser
                return NextResponse.redirect(new URL("/api/.ory/ui/login", request.url))
            default:
                console.log("default");
                return NextResponse.redirect(new URL("/api/.ory/ui/login", request.url))
        }
    })
}

function getRedirectUrl(newURL: string, request: NextRequest) {
    if (request.nextUrl.searchParams) {
        newURL = newURL + "?" + request.nextUrl.searchParams.toString()
    }

    const requestHeaders = new Headers(request.headers);

    if (request.nextUrl.pathname.startsWith('/oasis-api')) {
        requestHeaders.set('X-Openline-API-KEY', process.env.OASIS_API_KEY as string)
    } else if (request.nextUrl.pathname.startsWith('/customer-os-api')) {
        requestHeaders.set('X-Openline-API-KEY', process.env.CUSTOMER_OS_API_KEY as string)
    }

    return NextResponse.rewrite(new URL(newURL, request.url),
        {
            request: {
                headers: requestHeaders,
            },
        }
    )
}

export const config = {
    matcher: ['/oasis-api/(.*)', '/customer-os-api/(.*)'],
}
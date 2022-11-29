import { withAuth } from "next-auth/middleware"
import {NextRequest, NextResponse} from "next/server";

// More on how NextAuth.js middleware works: https://next-auth.js.org/configuration/nextjs#middleware

export default withAuth(function middleware(request: NextRequest) {
    var newURL=process.env.NEXT_PUBLIC_BE_PATH + "/" + request.nextUrl.pathname.substring(("/server/").length);
    console.log("Rewriting url to " + newURL);
    console.log("middleware: " + JSON.stringify(request.nextauth.token));

    const requestHeaders = new Headers(request.headers)
    requestHeaders.set('X-Openline-API-KEY', process.env.OASIS_API_KEY?process.env.OASIS_API_KEY:"")
    const response = NextResponse.next({
      request: {
        // New request headers
        headers: requestHeaders,
      },
    })
      return NextResponse.rewrite(new URL(newURL, request.url), 
        {
        request: {
          // New request headers
          headers: requestHeaders,
        },
      }
      )
  
}, 
{
  callbacks: {
    authorized({ req, token }) {
      console.log("Got Token: " + JSON.stringify(token));
      if(token) return true; // If there is a token, the user is authenticated
      return false;
    },
  },
})

export const config = {
  matcher: '/server/(.*)',
}
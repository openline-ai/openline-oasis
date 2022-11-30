import {Session} from "next-auth";

export async function loggedInOrRedirectToLogin(session: Session | null) {
    if (!session) {
        return {
            redirect: {
                destination: '/api/auth/signin',
                permanent: false,
            },
        }
    }

    return {
        props: { session }
    }
}
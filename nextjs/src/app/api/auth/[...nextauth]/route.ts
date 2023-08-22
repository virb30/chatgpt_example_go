import { AuthOptions } from 'next-auth';
import NextAuth from 'next-auth/next'
import KeycloakProvider from 'next-auth/providers/keycloak'

export const authConfig: AuthOptions = {
    providers: [
        KeycloakProvider({
            clientId: 'nextjs',
            clientSecret: '8RaPx3Nrq5z720nRdZV00Ea2diVyhHoF',
            issuer: 'http://host.docker.internal:8080/realms/master'
        })
    ]
}

const handler = NextAuth(authConfig);

export { handler as GET, handler as POST };
package wang.sorley.shiro.chapter2;

import org.apache.shiro.authc.AuthenticationException;
import org.apache.shiro.authc.AuthenticationInfo;
import org.apache.shiro.authc.AuthenticationToken;
import org.apache.shiro.authc.SimpleAuthenticationInfo;
import org.apache.shiro.authc.pam.AbstractAuthenticationStrategy;
import org.apache.shiro.realm.Realm;
import org.apache.shiro.util.CollectionUtils;

import java.util.Collection;

public class AtLeastTwoAuthenticationStrategy extends AbstractAuthenticationStrategy {

    @Override
    public AuthenticationInfo beforeAllAttempts(Collection<? extends Realm> realms, AuthenticationToken token) throws AuthenticationException {
        return new SimpleAuthenticationInfo();
    }

    @Override
    public AuthenticationInfo beforeAttempt(Realm realm, AuthenticationToken token, AuthenticationInfo aggregate) throws AuthenticationException {
        return aggregate;
    }

    @Override
    public AuthenticationInfo afterAttempt(Realm realm, AuthenticationToken token, AuthenticationInfo singleRealmInfo, AuthenticationInfo aggregateInfo, Throwable t) throws AuthenticationException {
        AuthenticationInfo authenticationInfo;
        if (singleRealmInfo == null) {
            authenticationInfo = aggregateInfo;
        } else if (aggregateInfo == null) {
            authenticationInfo = singleRealmInfo;
        } else {
            authenticationInfo = merge(singleRealmInfo, aggregateInfo);
        }
        return authenticationInfo;
    }

    @Override
    public AuthenticationInfo afterAllAttempts(AuthenticationToken token, AuthenticationInfo aggregate) throws AuthenticationException {
        if(aggregate == null || CollectionUtils.isEmpty(aggregate.getPrincipals()) || aggregate.getPrincipals().getRealmNames().size()<2){
            throw new AuthenticationException("Authentication token of type [" + token.getClass() + "]" +
                    "could not be authenticated by configured realms." + "please ensure that at least two realm can be athenticated these token");
        }
        return aggregate;
    }
}

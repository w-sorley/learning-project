package wang.sorley.shiro.chapter2;

import org.apache.shiro.authc.*;
import org.apache.shiro.realm.Realm;

public class CustomRealm2 implements Realm {

    public String getName() {
        return "CustomRealm2";
    }

    public boolean supports(AuthenticationToken authenticationToken) {
        return authenticationToken instanceof UsernamePasswordToken;
    }

    public AuthenticationInfo getAuthenticationInfo(AuthenticationToken authenticationToken) throws AuthenticationException {
        String username = (String) authenticationToken.getPrincipal();
        String passwrod = new String((char[]) authenticationToken.getCredentials());
        if (!"wang".equals(username)) {
            throw new UnknownAccountException("用户名错误");
        } else if (!"123".equals(passwrod)) {
            throw new IncorrectCredentialsException("密码错误");
        }
        return new SimpleAuthenticationInfo(username, passwrod, this.getName());
    }
}

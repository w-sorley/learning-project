package wang.sorley.shiro.chapter2;

import org.apache.shiro.authc.*;
import org.apache.shiro.realm.Realm;

public class CustomRealm3 implements Realm {
    public String getName() {
        return "CustomRealm3";
    }

    public boolean supports(AuthenticationToken authenticationToken) {
        return authenticationToken instanceof UsernamePasswordToken;
    }

    public AuthenticationInfo getAuthenticationInfo(AuthenticationToken authenticationToken) throws AuthenticationException {
        String username = (String) authenticationToken.getPrincipal();
        String password = new String((char[]) authenticationToken.getCredentials());
        if (!"zhang".equals(username)) {
            throw new UnknownAccountException("用户名错误");
        } else if (!"123".equals(password)) {
            throw new IncorrectCredentialsException("密码错误");
        }

        return new SimpleAuthenticationInfo(username + "@dtdream.com", password, this.getName());
    }
}

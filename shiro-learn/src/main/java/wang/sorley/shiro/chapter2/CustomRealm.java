package wang.sorley.shiro.chapter2;

import org.apache.shiro.authc.*;
import org.apache.shiro.realm.Realm;

/**
 * shiro学习1：自定义readlm
 * 重点:重写getAuthenticationInfo()方法，根据待认证的token返回认证信息
 * 过程:从待认证的token中获取必要信息，与数据源信息比对，不匹配返回响应的错误异常，匹配则返回info实例
 */
public class CustomRealm implements Realm {



    public String getName() {
        return "My Custom Realm 1";
    }
    public boolean supports(AuthenticationToken authenticationToken) {
        return authenticationToken instanceof UsernamePasswordToken;
    }

    public AuthenticationInfo getAuthenticationInfo(AuthenticationToken authenticationToken) throws AuthenticationException {
        String userName = (String)authenticationToken.getPrincipal();
        String passWord = new String((char[])authenticationToken.getCredentials());
        if(!"wang".equals(userName)){
            throw new UnknownAccountException("用户名错误");
        }else if(!"12358".equals(passWord)){
            throw new IncorrectCredentialsException("密码错误");
        }
        return new SimpleAuthenticationInfo(userName,passWord,this.getName());
    }
}

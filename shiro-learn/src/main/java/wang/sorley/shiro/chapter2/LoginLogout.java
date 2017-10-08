package wang.sorley.shiro.chapter2;

import org.apache.shiro.SecurityUtils;
import org.apache.shiro.authc.AuthenticationException;
import org.apache.shiro.authc.UsernamePasswordToken;
import org.apache.shiro.config.IniSecurityManagerFactory;
import org.apache.shiro.mgt.SecurityManager;
import org.apache.shiro.subject.Subject;
import org.apache.shiro.util.Factory;

/***
 * shiro学习1：简单登录/退出
 * 过程:SecurityManagerFactory -> SecurityManager ->SecurityUtils -> subject ->Token ->loggin -> logout
 */
public class LoginLogout {
    public static void main(String[] args) {
//        Factory<SecurityManager> securityManagerFactory = new IniSecurityManagerFactory("classpath:chapter2/shiro.ini");
//        Factory<SecurityManager> securityManagerFactory = new IniSecurityManagerFactory("classpath:chapter2/shiro-custom-realm.ini");  //测试自定义realm
        Factory<SecurityManager> securityManagerFactory = new IniSecurityManagerFactory("classpath:chapter2/shiro-jdbc-realm.ini");
        SecurityManager securityManager = securityManagerFactory.getInstance();
        SecurityUtils.setSecurityManager(securityManager);
        Subject subject = SecurityUtils.getSubject();
        UsernamePasswordToken token = new UsernamePasswordToken("wang","123");
        try {
            subject.login(token);
        }catch (AuthenticationException e){
            e.printStackTrace();
            System.out.println("身份验证失败！！");
        }

        if(subject.isAuthenticated()){
            System.out.println("身份验证成功");
        }
        subject.logout();
        if(!subject.isAuthenticated()){
            System.out.println("退出登录");
        }
    }
}

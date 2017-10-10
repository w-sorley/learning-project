package wang.sorley.shiro.chapter2;

import org.apache.shiro.SecurityUtils;
import org.apache.shiro.authc.AuthenticationException;
import org.apache.shiro.authc.UsernamePasswordToken;
import org.apache.shiro.config.IniSecurityManagerFactory;
import org.apache.shiro.mgt.SecurityManager;
import org.apache.shiro.subject.PrincipalCollection;
import org.apache.shiro.subject.Subject;
import org.apache.shiro.util.Factory;

import java.security.Principal;

/***
 * shiro学习1：简单登录/退出
 * 过程:SecurityManagerFactory -> SecurityManager ->SecurityUtils -> subject ->Token ->loggin -> logout
 */
public class LoginLogout {

    public static void main(String[] args) {
        String basicConfigFilePath = "classpath:chapter2/shiro.ini";
        String customRealmConfigFilePath = "classpath:chapter2/shiro-custom-realm.ini";
        String jdbcRealmConfigFileParh = "classpath:chapter2/shiro-jdbc-realm.ini";
        String authenticationAllConfigFilePath = "classpath:chapter2/authenticator-all-success.ini";
//        testLoginAndLogout(basicConfigFilePath);
//        testLoginAndLogout(customRealmConfigFilePath);
//        testLoginAndLogout(jdbcRealmConfigFileParh);
        testLoginAndLogout(authenticationAllConfigFilePath);


    }


    public static void testLoginAndLogout(String configFilePath) {
        Factory<SecurityManager> securityManagerFactory = new IniSecurityManagerFactory(configFilePath);
        SecurityManager securityManager = securityManagerFactory.getInstance();
        SecurityUtils.setSecurityManager(securityManager);
        Subject subject = SecurityUtils.getSubject();
        UsernamePasswordToken token = new UsernamePasswordToken("wang", "123");
        try {
            subject.login(token);
        } catch (AuthenticationException e) {
            e.printStackTrace();
            System.out.println("身份验证失败！！");
        }

        if (subject.isAuthenticated()) {
            System.out.println("身份验证成功");
        }
        PrincipalCollection principalCollection = subject.getPrincipals();
        System.out.println(principalCollection.asList().size());
        subject.logout();
        if (!subject.isAuthenticated()) {
            System.out.println("退出登录");
        }

    }


}

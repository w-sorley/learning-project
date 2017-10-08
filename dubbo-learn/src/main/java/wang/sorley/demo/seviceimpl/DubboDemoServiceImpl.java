package wang.sorley.demo.seviceimpl;

import wang.sorley.demo.service.DubboDemoService;

public class DubboDemoServiceImpl implements DubboDemoService {
    public String sayHello(String name) {
        return "Hello"+name;
    }
}

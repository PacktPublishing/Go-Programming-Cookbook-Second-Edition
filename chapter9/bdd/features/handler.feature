Feature: Bad Method
    Scenario: Good request
        Given we create a HandlerRequest payload with:
            | reader |
            | coder  |
            | other  |
        And we POST the HandlerRequest to /hello
        Then the response code should be 200
        And the response body should be:
            | BDD testing reader |
            | BDD testing coder  |
            | BDD testing other  |

# golang-microsoft-graph-rest-sample
Inspired by [active-directory-python-flask-graphapi-web-v2](https://github.com/Azure-Samples/active-directory-python-flask-graphapi-web-v2).

## Steps to Run
1. Register your Azure AD v2.0 app.  
    - Navigate to the [App Registration Portal](https://identity.microsoft.com). 
    - Go to the the `My Apps` page, click `Add an App`, and name your app.  
    - Set a platform by clicking `Add Platform`, select `Web`, and add a Redirect URI of ```http://localhost:5000/login/authorized```.
    - Click "Generate New Password" and record your Consumer Secret.  

2. Clone the code. 
    ```
    git clone https://github.com/Frederick-S/golang-microsoft-graph-rest-sample
    ```

3. In the top of main.go, add your Application/Client ID and Consumer Secret to the app config.

4. Run `go run main.go` in the terminal! Navigate to `http://localhost:5000`.

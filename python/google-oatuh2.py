from flask import Flask, redirect, request, session, url_for
import requests

app = Flask(__name__)
app.secret_key = 'YOUR_SECRET_KEY'

CLIENT_ID = 'YOUR_CLIENT_ID'
CLIENT_SECRET = 'YOUR_CLIENT_SECRET'
REDIRECT_URI = 'http://localhost:3000/oauth2callback'

@app.route('/')
def index():
    return 'Welcome to the Google OAuth 2.0 login demo.'

@app.route('/login')
def login():
    authorization_url = (
        'https://accounts.google.com/o/oauth2/v2/auth'
        '?response_type=code'
        '&client_id=' + CLIENT_ID +
        '&redirect_uri=' + REDIRECT_URI +
        '&scope=email%20profile'
    )
    return redirect(authorization_url)

@app.route('/oauth2callback')
def oauth2callback():
    code = request.args.get('code')
    token_url = 'https://oauth2.googleapis.com/token'
    payload = {
        'client_id': CLIENT_ID,
        'client_secret': CLIENT_SECRET,
        'code': code,
        'grant_type': 'authorization_code',
        'redirect_uri': REDIRECT_URI
    }
    response = requests.post(token_url, data=payload)
    tokens = response.json()
    session['access_token'] = tokens['access_token']
    session['refresh_token'] = tokens.get('refresh_token')
    return redirect(url_for('profile'))

@app.route('/profile')
def profile():
    access_token = session.get('access_token')
    userinfo_url = 'https://www.googleapis.com/oauth2/v1/userinfo'
    headers = {'Authorization': f'Bearer {access_token}'}
    response = requests.get(userinfo_url, headers=headers)
    user_info = response.json()
    return f'Hello, {user_info["name"]}!'

if __name__ == '__main__':
    app.run(port=3000)


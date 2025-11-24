from flask import Flask, request, jsonify

app = Flask(__name__)

# 1. API1: Broken Object Level Authorization
@app.route('/api1/items/<item_id>', methods=['GET'])
def api1(item_id):
    # 漏洞点：未做对象级别权限校验
    return jsonify({'item_id': item_id, 'detail': 'Object info (no auth check)'})

# 2. API2: Broken User Authentication
@app.route('/api2/login', methods=['POST'])
def api2():
    # 漏洞点：弱认证逻辑
    data = request.json
    if data and data.get('username') == 'admin' and data.get('password') == '123456':
        return jsonify({'msg': 'Login success', 'token': 'fake-jwt-token'})
    return jsonify({'msg': 'Login failed'}), 401

# 3. API3: Excessive Data Exposure
@app.route('/api3/userinfo', methods=['GET'])
def api3():
    # 漏洞点：返回敏感字段
    return jsonify({'username': 'alice', 'email': 'alice@example.com', 'password': 'plaintextpassword'})

# 4. API4: Lack of Resources & Rate Limiting
@app.route('/api4/nolimit', methods=['GET'])
def api4():
    # 漏洞点：无速率/资源限制
    return jsonify({'msg': 'No rate limit here'})

# 5. API5: Broken Function Level Authorization
@app.route('/api5/admin', methods=['GET'])
def api5():
    # 漏洞点：无功能级别权限校验
    return jsonify({'msg': 'Admin function accessed'})

# 6. API6: Mass Assignment
@app.route('/api6/profile', methods=['POST'])
def api6():
    # 漏洞点：允许批量赋值
    profile = request.json
    return jsonify({'msg': 'Profile updated', 'profile': profile})

# 7. API7: Security Misconfiguration
@app.route('/api7/debug', methods=['GET'])
def api7():
    # 漏洞点：暴露调试信息
    return jsonify({'debug': True, 'config': 'Sensitive config here'})

# 8. API8: Injection
@app.route('/api8/search', methods=['POST'])
def api8():
    # 漏洞点：未过滤输入，模拟SQL注入
    query = request.json.get('q', '')
    return jsonify({'result': f"You searched for: {query}"})

# 9. API9: Improper Assets Management
@app.route('/api9/old-api', methods=['GET'])
def api9():
    # 漏洞点：暴露废弃API
    return jsonify({'msg': 'Deprecated API still accessible'})

# 10. API10: Insufficient Logging & Monitoring
@app.route('/api10/transfer', methods=['POST'])
def api10():
    # 漏洞点：无日志记录
    data = request.json
    return jsonify({'msg': 'Transfer completed', 'data': data})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001, debug=True)

### 登陆接口

**功能描述**: 处理用户登录请求，验证用户名和密码，成功时返回一个token。

**请求方法**: `POST`

**URL**: `/api/logIn`

**请求参数**:

```json
{
  "username": "string",
  "password": "string"
}
```

**成功响应**:

- **Code**: `200`
- **Content**:

```json
{
  "msg": "yeah",
  "token": "string"
}
```

**错误响应**:

- 数据错误:

  - **Code**: `200`
  - **Content**:

  ```json
  {
    "msg": "data error"
  }
  ```

- 空输入:

  - **Code**: `200`
  - **Content**:

  ```json
  {
    "msg": "empty enter"
  }
  ```

- 数据库错误:

  - **Code**: `200`
  - **Content**:

  ```json
  {
    "msg": "database error"
  }
  ```

- 账户不存在:

  - **Code**: `200`
  - **Content**:

  ```json
  {
    "msg": "no this account"
  }
  ```

  

### 注册接口

**功能描述**: 处理用户注册请求，创建新用户账户，并返回一个token。

**请求方法**: `POST`

**URL**: `/api/logUp`

**请求参数**:

```json
{
  "username": "string",
  "password": "string",
  "email": "string"
}
```

**成功响应**:

- **Code**: `200`
- **Content**:

```json
{
  "msg": "yeah",
  "token": "string"
}
```

**错误响应**:

- 数据错误:

  - **Code**: `200`
  - **Content**:

  ```json
  {
    "msg": "data error"
  }
  ```

- 空输入:

  - **Code**: `200`
  - **Content**:

  ```json
  {
    "msg": "empty enter"
  }
  ```

- 数据库错误:

  - **Code**: `200`
  - **Content**:

  ```json
  {
    "msg": "database error"
  }
  ```

- 用户已存在:

  - **Code**: `200`
  - **Content**:

  ```json
  {
    "msg": "user already existed"
  }
  ```

  

### 列表接口

**功能描述**: 获取用户视频资源列表。

**请求方法**: `GET`

**URL**: `/api/getList`

**请求头**:

- **Authorization**: `token`

**成功响应**:

- **Code**: `200`
- **Content**:

```json
{
  "files": ["array", "of", "file", "names"]
}
```

**错误响应**:

- 键不存在或已过期:

  - **Code**: `200`
  - **Content**:

  ```json
  {
    "msg": "token不存在或已过期"
  }
  ```

- 服务器错误 (例如无法打开目录):

  - **Code**: `500`
  - **Content**:

  ```json
  {
    "msg": "Internal Server Error"
  }
  ```

**备注**:

- 用户需要提供有效的token来验证身份。

- 接口将返回与用户下的所有文件名。

  

  ### 获取视频接口

  **功能描述**: 根据提供的文件名获取指定的视频资源。

  **请求方法**: `POST`

  **URL**: `/api/getVideo`

  **请求头**:

  - **Authorization**: `token`

  **请求参数**:

  ```json
  {
    "filename": "string"
  }
  ```

  **成功响应**:

  - **Code**: `200`
  - **Content**: 返回请求的视频文件。

  **错误响应**:

  - token不存在或已过期:

    - **Code**: `200`
    - **Content**:

    ```json
    {
      "msg": "token不存在或已过期"
    }
    ```

  **备注**:

  - 用户需要提供有效的token来验证身份。

  - 请求参数需要在请求体中以JSON格式发送。

  - 接口将根据提供的文件名返回指定的视频文件。

    

    ### Upload 接口

    **功能描述**: 处理multipart表单的POST请求，允许用户上传视频文件到他们的账号目录。

    **请求方法**: `POST`

    **URL**: `/api/upLoad`

    **请求头**:

    - **Authorization**: `token`

    **表单数据**:

    - **file**: 用户需要上传的视频文件。

    **成功响应**:

    - **Code**: `200`
    - **Content**:

    ```json
    {
      "message": "文件上传成功"
    }
    ```

    **错误响应**:

    - token不存在或已过期:

      - **Code**: `200`
      - **Content**:

      ```json
      {
        "msg": "token不存在或已过期"
      }
      ```

    - 服务器错误 (例如无法获取表单文件):

      - **Code**: `500`
      - **Content**:

      ```json
      {
        "error": "string"
      }
      ```

    **备注**:

    - 用户需要提供有效的token来验证身份。
    - 如果文件上传成功，用户将响应确认消息。

    
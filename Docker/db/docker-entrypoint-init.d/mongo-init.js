var users = [
    {
      user: '****',
      pwd: '****',
      roles: [
        {
          role: 'dbOwner',
          db: '****'
        }
      ]
    }
  ];
  
  for (var i = 0, length = users.length; i < length; ++i) {
    db.createUser(users[i]);
  }
# Know Your Customer System

## 1. Stack

### Backend
- Golang
- Gin Gonic
- Swaggo / swag
- Lumberjack
- GORM (PostgreSQL driver: pq)

### Database
- PostgreSQL

### Infrastructure
- Docker
- Redis - IN PROCCESS 
- Kafka - IN PROCCESS

## 2. Endpoints

### /auth/sign-up POST endpoint

- Accepts login       | password    | 
- returns models.User | status_code |

### /auth/sign-in POST endpoint

- Accepts login        | password   |
- returns access_token | token_type |

### /profile/get GET endpoint

- Accepts headers access_token |             |
- returns models.Profile       | status_code |

### /profile/update PUT endpoint

- Accepts headers access_token | models.Profile |
- returns models.Profile       | status_code    |

### /documents/update PUT endpoint

- Accepts headers access_token | file.frontID | file.backID | file.selfieID |
- returns models.Documents     | status_code  |             |               |

### /upload GET endpoint

- Accepts filePath |             |
- returns file     | status_code |

### /user/info GET endpoint

- Accepts headers access_token |             |
- returns models.User          | status_code |

### /admin/users GET endpoint

- Accepts headers access_token | filters_status | filters_user_id | 
- returns models.User          | models.Profile | status_code     |

### /admin/user-docs GET endpoint

- Accepts headers access_token | user_id     | 
- returns file: str(filePath)  | status_code |

### /admin/confirm POST endpoint 

- Accepts headers access_token | user_id     |
- returns models.User          | status_code |
# Couply

A dating app concept by Tanricko

## GitHub Repository

<https://github.com/wendao2000/couply>

## Functional Requirements

### 1. Authentication

- Sign Up
  - Email registration
  - SSO (Google, Apple, etc.)
- Login
  - Email/password
  - Phone/OTP
  - SSO
- Password management (reset, change)

### 2. User Management

- Profile CRUD
  - Basic info
  - Photos
  - Preferences
- Account roles
  - Free user (10 swipes/day)
  - Premium user (unlimited swipes)

### 3. Matching System

- Swipe functionality
  - Like/Dislike recording
  - Swipe limit enforcement
  - 24-hour profile cooldown
- Matching algorithm
  - Real-time match generation
  - Preference-based filtering
  - Learning from user actions

### 4. Chat System

- Match-only chat rooms
- End-to-end encryption
- Message history
- Online status

### 5. Premium Features

- Payment processing
- Role upgrade
- Feature unlocking

### 6. Support System

- Help center/FAQ
- Contact support
- Privacy/Terms pages

## Non-Functional Requirements

### 1. Performance

- API latency: <200ms for critical sync endpoints (Auth, Chat, Payment)
- Cache hit ratio: >80%
- Retry mechanism: exponential backoff with max retries

### 2. Security

- Password: bcrypt + salt
- JWT auth + session management
- HTTPS/TLS encryption
- XSS/CSRF protection
- Rate limiting

### 3. Scalability

- Microservice architecture
- Horizontal scaling capability
- Auto-scaling triggers
- Load balancing
- DB sharding strategy

### 4. Availability

- 99.9% uptime SLA
- Automated failover
- Regular backups
- Disaster recovery plan

### 5. Monitoring

- System metrics & health checks
- Error tracking & logging
- Resource utilization alerts
- API performance tracking

### 6. Compliance

- Local privacy laws
- Payment security
- Data retention policy
- Age verification

## Tech Stack

### Backend Language: Go (Golang)

- I'm very experienced with it
- Built-in tools make concurrent processing simple
- Easy to build and scale services

### Primary Database: PostgreSQL (Relational Database)

- Stores all user data and matches
- Fixed structure enforce data consistency across users
- ACID compliance ensures data reliability and transaction safety

### Caching Layer: Redis

- Stores frequently accessed data in memory
- Reduce database load
- Helps track daily swipe counts efficiently

### Message Queue: RabbitMQ

- Processes matches asynchronously - users can keep swiping
- Queues tasks like match processing and notifications
- Guarantees message delivery with acknowledgments (prevents lost messages)

### Real-Time Communication

#### Server-Sent Events (SSE)

- Sends match and swiping limit notifications to users' device

#### WebSocket

- Powers the real-time chat feature

#### Push Notifications

- Sends notifications to users' phones

### Object Storage: Amazon S3

- Stores all user photos and files
- Easy to manage and scale storage
- Cost effective for large amounts of files

### CDN: Akamai

- Stores copies of photos in multiple locations worldwide for faster access

### Version Control: GitHub

- Stores and tracks all code changes
- Most widely used platform for code management

### CI/CD: GitHub Actions

- Automatically tests and deploys code
- Built into GitHub, easy to set up
- Use our own servers to run it, saves money in the long run

## Entity Relationship Diagram

![Couply ERD](couply_erd.png)

## Sequence Diagram

### Swiping

![Couply SQD - Swiping](couply_sqd_swiping.png)

### Payment

![Couply SQD - Payment](couply_sqd_payment.png)

## Test Cases

| Scenario | Services Involved | Test Case | Description | Action | Expected Result |
|----------|-------------------|------------|-------------|---------|-----------------|
| Registration | - Auth Service<br>- Profile Service<br>- Notifier Service | New User Setup | Complete user registration and profile creation | 1. Submit registration details<br>2. Verify email<br>3. Complete initial profile<br>4. Add profile photos | 1. Account created successfully<br>2. Email verified<br>3. Profile saved<br>4. Photos uploaded and visible |
| Profile Updates | - Profile Service<br>- Matchmaker Service | Profile Customization | User updates preferences and photos | 1. Update user preferences<br>2. Change profile photos<br>3. Edit bio and details<br>4. Check profile visibility | 1. Changes saved correctly<br>2. Photos processed and ordered<br>3. Profile appears in recommendations |
| Free User Swiping | - Matchmaker Service<br>- Profile Service<br>- Notifier Service | Daily Swipe Limit | Free user reaches daily swipe limit | 1. Make 10 swipes<br>2. Attempt 11th swipe<br>3. Check limit notification | 1. First 10 swipes succeed<br>2. 11th swipe blocked<br>3. Premium upgrade prompt shown |
| Premium Upgrade | - Payment Service<br>- Profile Service<br>- Matchmaker Service<br>- Notifier Service | Premium Conversion | User upgrades to premium account | 1. Start premium purchase<br>2. Complete payment<br>3. Test premium features<br>4. Continue swiping | 1. Payment successful<br>2. Premium status active<br>3. Swipe limit removed<br>4. All premium features available |
| Match Creation | - Matchmaker Service<br>- Chat Service<br>- Notifier Service | Mutual Match | Two users match with each other | 1. User A likes User B<br>2. User B likes User A<br>3. Check match creation | 1. Match recorded<br>2. Chat room created<br>3. Both users notified<br>4. Users appear in each other's matches |
| Chat Interaction | - Chat Service<br>- Profile Service<br>- Notifier Service | Active Chat | Users engage in conversation | 1. Open chat room<br>2. Exchange messages<br>3. Close and reopen chat<br>4. Check notifications | 1. Messages deliver instantly<br>2. Chat history preserved<br>3. Notifications work properly |
| Unmatch Process | - Chat Service<br>- Matchmaker Service<br>- Notifier Service | User Unmatches | User decides to unmatch | 1. Initiate unmatch<br>2. Confirm action<br>3. Check both users' states | 1. Match removed<br>2. Chat room archived<br>3. Users can't message<br>4. Neither sees the match |

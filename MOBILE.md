# Mobile App Plan

This document describes how to build a mobile app for School Awesome using free tools, and how to host it or run it without cost for development.

## Why mobile?

The app can reach students, parents, and teachers on the go.
It should reuse the backend API and focus on a mobile-friendly experience.

## Recommended approach

### Use Expo + React Native

Expo is the easiest free option for building mobile apps.
It works with the existing React knowledge and enables running on real devices through Expo Go.

### Why Expo?

- Free to start
- No native build tooling required for development
- Supports Android and iOS
- Easy to test with Expo Go app
- Good fit with React-based UI logic

## Mobile app structure

Create a new folder or repo for the mobile app:

```bash
npx create-expo-app mobile
```

Recommended structure:

- `mobile/App.js` — app entrypoint
- `mobile/screens/LoginScreen.js`
- `mobile/screens/HomeScreen.js`
- `mobile/screens/ProfileScreen.js`
- `mobile/screens/UsersScreen.js`
- `mobile/services/api.js`
- `mobile/components/ProtectedRoute.js`

## API usage

Reuse the backend API endpoints:
- `POST /api/v1/auth/login`
- `GET /api/v1/me`
- `GET /api/v1/users`
- `POST /api/v1/admin/students`
- `POST /api/v1/admin/teachers`

Use a shared API client module to manage access tokens.

## Free app hosting / distribution

### Option 1: Expo Go (free)

- Install Expo CLI
- Run `npx expo start`
- Scan the QR code with Expo Go on Android or iOS
- This is free and ideal for development

### Option 2: Publish with Expo's free hosting

Expo offers a free publish URL for development builds:

```bash
npx expo publish
```

This creates a URL your team can open inside Expo Go.

### Option 3: Web version of Expo app

Expo can build a web app too.
You can host the web output for free on Vercel or Netlify.

```bash
npx expo build:web
```

### Option 4: PWA (quick mobile web)

If you want the fastest route, make the existing web app mobile-friendly and host it on:
- Netlify
- Vercel
- GitHub Pages
- Firebase Hosting

A PWA is not a native app, but it works on mobile and is free to host.

## Recommended free hosting flow

1. Host backend on Render or Railway free tier.
2. Host frontend on Vercel / Netlify free tier.
3. Build an Expo app and use Expo Go for mobile testing.
4. Optionally publish the mobile app URL via Expo.

## Getting started with Expo

1. Install Expo CLI:

```bash
npm install -g expo-cli
```

2. Create the app:

```bash
npx create-expo-app mobile
```

3. Start development:

```bash
cd mobile
npx expo start
```

4. Open on device with Expo Go.

## Suggested first mobile screens

- Login
- Dashboard / home
- Profile
- User list
- Notifications
- Settings

## Notes

- Use API tokens for authentication.
- Build simple mobile UIs first, then add modules.
- Start with login + profile + list screens.
- Use the same backend API routes as the web app.
- Save app state locally using SecureStore or AsyncStorage.

## Next mobile milestones

- [ ] Mobile login and profile
- [ ] Class/Student list view
- [ ] Admin create student/teacher
- [ ] Notifications screen
- [ ] Homework/fees quick links

## Free deployment summary

- Backend: free platform with Docker support (Render / Railway)
- Frontend: Vercel / Netlify
- Mobile dev: Expo Go
- Web PWA: Vercel / Netlify

# School Awesome Feature Roadmap

This document captures planned modules, future direction, and a daily build approach for the School Awesome application.

## Vision

School Awesome should evolve into a full-featured school management platform that supports:
- Web administration and teacher access
- Future mobile application access for students, parents, and staff
- Modular expansion with clearly defined feature domains
- A modern, colorful, responsive UI

## Core future modules

### 1. Role Management

Provide a centralized role and permission management module:
- Define roles such as `admin`, `principal`, `teacher`, `student`, `accountant`, `librarian`, `parent`
- Assign users to roles and role-based access levels
- Configure module-level permissions:
  - who can view student data
  - who can manage fees
  - who can publish homework
  - who can send notifications
- Support custom role creation for organization-specific staff structures

### 2. Class Management

Organize students by class and section:
- Create classes, sections, and academic years
- Assign teachers to classes
- Enroll students into classes and sections
- Display class rosters, teacher assignments, and schedules

### 3. Student Management

Build a student-centric module:
- Student profiles with academic history
- Class-based listing and filtering
- Attendance and performance metadata
- Student dashboard for personal progress

### 4. Library Management

Manage library operations:
- Books / inventory catalog
- Issue and return tracking
- Late fee calculations
- Librarian role with access controls
- Search and book reservation flow

### 5. Learning Management System (LMS)

Add digital learning capabilities:
- Course and lesson creation
- Assignment upload and grading
- Resource sharing for students and teachers
- Class-specific learning modules

### 6. Fees Management

Handle finance workflows:
- Fee categories and instalment plans
- Student fee schedules and invoices
- Payment tracking and receipts
- Alerts for overdue fees
- Accountant dashboard for revenue summary

### 7. Homework & Assignments

Support academic task management:
- Post homework by class/section
- Attach files or links
- Deadline and status tracking
- Teacher grading and feedback
- Student submission interface

### 8. Notification System

Send notifications across the school:
- Broadcast to students, parents, staff, or selected classes
- Delivery via email, in-app notifications, or SMS (future)
- Notification history and templates
- Role-based notification permissions

### 9. Reporting & Analytics

Provide operational and academic reports:
- Attendance reports
- Fee collection and outstanding balances
- Student performance summaries
- Class growth metrics
- Role-based dashboards for admins and teachers

### 10. Mobile Application Support

Plan for a mobile-friendly extension of the platform:
- Mobile-first UI design from the start
- API endpoints designed for mobile clients
- Separate mobile app modules for:
  - student view
  - parent view
  - teacher tasks
  - staff notifications
- Future mobile features:
  - real-time notifications
  - on-the-go attendance
  - mobile fee payments
  - homework submission

## UI and Design Goals

- Build a modern, colorful, responsive interface
- Use clear role-specific navigation
- Support dark/light theme in future releases
- Keep screens minimal and action-driven
- Use visual cards, badges, and progress indicators for school data

## Daily Build / Module Plan

Use a daily incremental approach for module development:
1. Day 1: Role management base + permission model
2. Day 2: Class management and student grouping
3. Day 3: Student profile and class roster UI
4. Day 4: Library module scaffold and book catalog data model
5. Day 5: Fees module scaffold and payment tracking
6. Day 6: Homework module scaffold and assignment flow
7. Day 7: Notification module scaffold and admin broadcast
8. Day 8: Reporting module scaffold and basic dashboard
9. Day 9: Mobile API readiness and responsive UI updates
10. Day 10+: polish UI, integrate modules, add analytics

## Suggested repo structure for modules

- `internal/core/domain` — domain models for new modules
- `internal/core/usecase` — business logic for each module
- `internal/adapter/api` — HTTP handlers for module endpoints
- `frontend/src/pages` — separate pages for module UIs
- `frontend/src/components` — reusable UI components
- `frontend/src/services` — API client wrappers

## Notes

- Treat each module as an independent feature slice.
- Keep API contracts stable for both web and future mobile clients.
- Start with role and class systems first, because they are foundational.
- Use the existing clean architecture pattern for new modules.
- Track module progress with daily TODOs or tickets.

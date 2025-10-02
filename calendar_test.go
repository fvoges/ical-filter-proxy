package main

import (
	"testing"

	ics "github.com/arran4/golang-ical"
)

func TestAnonymizeEvent(t *testing.T) {
	// Create a test event with various properties
	cal := ics.NewCalendar()
	event := cal.AddEvent("test-event-123")

	// Set various properties that should be removed
	event.SetSummary("Confidential Meeting")
	event.SetDescription("Secret discussion about project X")
	event.SetLocation("Conference Room A")
	event.SetURL("https://zoom.us/j/123456789")
	event.SetOrganizer("organizer@example.com")
	event.AddAttendee("attendee1@example.com")
	event.AddAttendee("attendee2@example.com")

	// Anonymize the event
	AnonymizeEvent(event)

	// Verify summary is set to "Busy"
	summary := event.GetProperty(ics.ComponentPropertySummary)
	if summary == nil || summary.Value != "Busy" {
		t.Errorf("Expected summary to be 'Busy', got '%v'", summary)
	}

	// Verify description is cleared
	description := event.GetProperty(ics.ComponentPropertyDescription)
	if description != nil && description.Value != "" {
		t.Errorf("Expected description to be empty, got '%s'", description.Value)
	}

	// Verify location is cleared
	location := event.GetProperty(ics.ComponentPropertyLocation)
	if location != nil && location.Value != "" {
		t.Errorf("Expected location to be empty, got '%s'", location.Value)
	}

	// Verify URL is cleared
	url := event.GetProperty(ics.ComponentPropertyUrl)
	if url != nil && url.Value != "" {
		t.Errorf("Expected URL to be empty, got '%s'", url.Value)
	}

	// Verify components (alarms) are removed
	if len(event.Components) > 0 {
		t.Errorf("Expected components to be removed, got %d components", len(event.Components))
	}
}

func TestStringMatchRule_hasConditions(t *testing.T) {
	tests := []struct {
		name     string
		rule     StringMatchRule
		expected bool
	}{
		{
			name:     "empty rule",
			rule:     StringMatchRule{},
			expected: false,
		},
		{
			name:     "null condition",
			rule:     StringMatchRule{Null: true},
			expected: true,
		},
		{
			name:     "contains condition",
			rule:     StringMatchRule{Contains: "test"},
			expected: true,
		},
		{
			name:     "prefix condition",
			rule:     StringMatchRule{Prefix: "test"},
			expected: true,
		},
		{
			name:     "suffix condition",
			rule:     StringMatchRule{Suffix: "test"},
			expected: true,
		},
		{
			name:     "regex condition",
			rule:     StringMatchRule{RegexMatch: "test.*"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.rule.hasConditions()
			if result != tt.expected {
				t.Errorf("hasConditions() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestStringMatchRule_matchesString(t *testing.T) {
	tests := []struct {
		name     string
		rule     StringMatchRule
		data     string
		expected bool
	}{
		{
			name:     "null match on empty string",
			rule:     StringMatchRule{Null: true},
			data:     "",
			expected: true,
		},
		{
			name:     "null no match on non-empty string",
			rule:     StringMatchRule{Null: true},
			data:     "test",
			expected: false,
		},
		{
			name:     "contains match",
			rule:     StringMatchRule{Contains: "meeting"},
			data:     "Team meeting at 2pm",
			expected: true,
		},
		{
			name:     "contains no match",
			rule:     StringMatchRule{Contains: "meeting"},
			data:     "Conference call",
			expected: false,
		},
		{
			name:     "prefix match",
			rule:     StringMatchRule{Prefix: "Team"},
			data:     "Team meeting",
			expected: true,
		},
		{
			name:     "prefix no match",
			rule:     StringMatchRule{Prefix: "Team"},
			data:     "Daily meeting",
			expected: false,
		},
		{
			name:     "suffix match",
			rule:     StringMatchRule{Suffix: "meeting"},
			data:     "Team meeting",
			expected: true,
		},
		{
			name:     "suffix no match",
			rule:     StringMatchRule{Suffix: "meeting"},
			data:     "Team call",
			expected: false,
		},
		{
			name:     "regex match",
			rule:     StringMatchRule{RegexMatch: "^[0-9]{3}$"},
			data:     "123",
			expected: true,
		},
		{
			name:     "regex no match",
			rule:     StringMatchRule{RegexMatch: "^[0-9]{3}$"},
			data:     "abc",
			expected: false,
		},
		{
			name:     "multiple conditions all match",
			rule:     StringMatchRule{Contains: "meeting", Prefix: "Team"},
			data:     "Team meeting",
			expected: true,
		},
		{
			name:     "multiple conditions one fails",
			rule:     StringMatchRule{Contains: "meeting", Prefix: "Team"},
			data:     "Daily meeting",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.rule.matchesString(tt.data)
			if result != tt.expected {
				t.Errorf("matchesString() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestFilter_matchesEvent(t *testing.T) {
	// Create a test calendar and event
	cal := ics.NewCalendar()
	event := cal.AddEvent("test-event")
	event.SetSummary("Team Meeting")
	event.SetDescription("Weekly sync")
	event.SetLocation("Room 101")

	tests := []struct {
		name     string
		filter   Filter
		expected bool
	}{
		{
			name: "match on summary contains",
			filter: Filter{
				Match: EventMatchRules{
					Summary: StringMatchRule{Contains: "Meeting"},
				},
			},
			expected: true,
		},
		{
			name: "no match on summary contains",
			filter: Filter{
				Match: EventMatchRules{
					Summary: StringMatchRule{Contains: "Conference"},
				},
			},
			expected: false,
		},
		{
			name: "match on description contains",
			filter: Filter{
				Match: EventMatchRules{
					Description: StringMatchRule{Contains: "sync"},
				},
			},
			expected: true,
		},
		{
			name: "match on location prefix",
			filter: Filter{
				Match: EventMatchRules{
					Location: StringMatchRule{Prefix: "Room"},
				},
			},
			expected: true,
		},
		{
			name: "match multiple conditions",
			filter: Filter{
				Match: EventMatchRules{
					Summary:  StringMatchRule{Contains: "Meeting"},
					Location: StringMatchRule{Prefix: "Room"},
				},
			},
			expected: true,
		},
		{
			name: "no match on multiple conditions (one fails)",
			filter: Filter{
				Match: EventMatchRules{
					Summary:  StringMatchRule{Contains: "Meeting"},
					Location: StringMatchRule{Prefix: "Building"},
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.filter.matchesEvent(*event)
			if result != tt.expected {
				t.Errorf("matchesEvent() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

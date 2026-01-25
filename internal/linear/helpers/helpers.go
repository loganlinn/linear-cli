// Package helpers provides backward-compatible wrappers for refactored packages.
//
// DEPRECATED: This package is maintained for backward compatibility only.
// New code should import the specific packages directly:
//   - github.com/joa23/linear-cli/internal/linear/identifiers
//   - github.com/joa23/linear-cli/internal/linear/guidance
//   - github.com/joa23/linear-cli/internal/linear/validation
//   - github.com/joa23/linear-cli/internal/linear/metadata
//   - github.com/joa23/linear-cli/internal/linear/pagination
package helpers

import (
	"github.com/joa23/linear-cli/internal/linear/core"
	"github.com/joa23/linear-cli/internal/linear/guidance"
	"github.com/joa23/linear-cli/internal/linear/identifiers"
	"github.com/joa23/linear-cli/internal/linear/metadata"
	"github.com/joa23/linear-cli/internal/linear/pagination"
	"github.com/joa23/linear-cli/internal/linear/validation"
)

// Validation constants (from validation package)
const (
	// Deprecated: Use validation.MaxTitleLength instead
	MaxTitleLength = validation.MaxTitleLength
	// Deprecated: Use validation.MaxDescriptionLength instead
	MaxDescriptionLength = validation.MaxDescriptionLength
	// Deprecated: Use validation.MaxNotificationLimit instead
	MaxNotificationLimit = validation.MaxNotificationLimit
)

// Identifier detection functions (from identifiers package)

// Deprecated: Use identifiers.IsIssueIdentifier instead
func IsIssueIdentifier(s string) bool {
	return identifiers.IsIssueIdentifier(s)
}

// Deprecated: Use identifiers.IsEmail instead
func IsEmail(s string) bool {
	return identifiers.IsEmail(s)
}

// Deprecated: Use identifiers.IsUUID instead
func IsUUID(s string) bool {
	return identifiers.IsUUID(s)
}

// Error guidance functions (from guidance package)

// Deprecated: Use guidance.ValidationErrorWithExample instead
func ValidationErrorWithExample(field, requirement, example string) error {
	return guidance.ValidationErrorWithExample(field, requirement, example)
}

// Deprecated: Use guidance.OperationFailedError instead
func OperationFailedError(operation, resourceType string, guidanceList []string) error {
	return guidance.OperationFailedError(operation, resourceType, guidanceList)
}

// Deprecated: Use guidance.EnhanceGenericError instead
func EnhanceGenericError(operation string, err error) error {
	return guidance.EnhanceGenericError(operation, err)
}

// Deprecated: Use guidance.InvalidStateIDError instead
func InvalidStateIDError(stateID string, err error) error {
	return guidance.InvalidStateIDError(stateID, err)
}

// Deprecated: Use guidance.ResourceNotFoundError instead
func ResourceNotFoundError(resourceType, resourceID string, err error) error {
	return guidance.ResourceNotFoundError(resourceType, resourceID, err)
}

// Validation functions (from validation package)

// Deprecated: Use validation.IsValidEmoji instead
func IsValidEmoji(emoji string) bool {
	return validation.IsValidEmoji(emoji)
}

// Deprecated: Use validation.IsValidMetadataKey instead
func IsValidMetadataKey(key string) bool {
	return validation.IsValidMetadataKey(key)
}

// Deprecated: Use validation.ValidateStringLength instead
func ValidateStringLength(value, fieldName string, maxLength int) error {
	return validation.ValidateStringLength(value, fieldName, maxLength)
}

// Deprecated: Use validation.ValidatePositiveIntWithRange instead
func ValidatePositiveIntWithRange(value int, fieldName string, min, max int) error {
	return validation.ValidatePositiveIntWithRange(value, fieldName, min, max)
}

// Metadata functions (from metadata package)

// Deprecated: Use metadata.ExtractMetadataFromDescription instead
func ExtractMetadataFromDescription(description string) (map[string]interface{}, string) {
	return metadata.ExtractMetadataFromDescription(description)
}

// Deprecated: Use metadata.InjectMetadataIntoDescription instead
func InjectMetadataIntoDescription(description string, metadataMap map[string]interface{}) string {
	return metadata.InjectMetadataIntoDescription(description, metadataMap)
}

// Deprecated: Use metadata.UpdateDescriptionPreservingMetadata instead
func UpdateDescriptionPreservingMetadata(oldDescription, newDescription string) string {
	return metadata.UpdateDescriptionPreservingMetadata(oldDescription, newDescription)
}

// Pagination functions (from pagination package)

// Deprecated: Use pagination.ValidatePagination instead
func ValidatePagination(input *core.PaginationInput) *core.PaginationInput {
	return pagination.ValidatePagination(input)
}

// Deprecated: Use pagination.MapSortField instead
func MapSortField(sort string) string {
	return pagination.MapSortField(sort)
}

// Deprecated: Use pagination.MapSortDirection instead
func MapSortDirection(direction string) string {
	return pagination.MapSortDirection(direction)
}

// ErrorWithGuidance is a deprecated type, kept for compatibility
// Deprecated: Use errors from the guidance package instead
type ErrorWithGuidance = guidance.ErrorWithGuidance

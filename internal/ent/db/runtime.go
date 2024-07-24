// Code generated by ent, DO NOT EDIT.

package db

import (
	"time"

	"github.com/openmeterio/openmeter/internal/ent/db/balancesnapshot"
	"github.com/openmeterio/openmeter/internal/ent/db/entitlement"
	"github.com/openmeterio/openmeter/internal/ent/db/feature"
	"github.com/openmeterio/openmeter/internal/ent/db/grant"
	"github.com/openmeterio/openmeter/internal/ent/db/usagereset"
	"github.com/openmeterio/openmeter/internal/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	balancesnapshotMixin := schema.BalanceSnapshot{}.Mixin()
	balancesnapshotMixinFields0 := balancesnapshotMixin[0].Fields()
	_ = balancesnapshotMixinFields0
	balancesnapshotMixinFields1 := balancesnapshotMixin[1].Fields()
	_ = balancesnapshotMixinFields1
	balancesnapshotFields := schema.BalanceSnapshot{}.Fields()
	_ = balancesnapshotFields
	// balancesnapshotDescNamespace is the schema descriptor for namespace field.
	balancesnapshotDescNamespace := balancesnapshotMixinFields0[0].Descriptor()
	// balancesnapshot.NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	balancesnapshot.NamespaceValidator = balancesnapshotDescNamespace.Validators[0].(func(string) error)
	// balancesnapshotDescCreatedAt is the schema descriptor for created_at field.
	balancesnapshotDescCreatedAt := balancesnapshotMixinFields1[0].Descriptor()
	// balancesnapshot.DefaultCreatedAt holds the default value on creation for the created_at field.
	balancesnapshot.DefaultCreatedAt = balancesnapshotDescCreatedAt.Default.(func() time.Time)
	// balancesnapshotDescUpdatedAt is the schema descriptor for updated_at field.
	balancesnapshotDescUpdatedAt := balancesnapshotMixinFields1[1].Descriptor()
	// balancesnapshot.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	balancesnapshot.DefaultUpdatedAt = balancesnapshotDescUpdatedAt.Default.(func() time.Time)
	// balancesnapshot.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	balancesnapshot.UpdateDefaultUpdatedAt = balancesnapshotDescUpdatedAt.UpdateDefault.(func() time.Time)
	entitlementMixin := schema.Entitlement{}.Mixin()
	entitlementMixinFields0 := entitlementMixin[0].Fields()
	_ = entitlementMixinFields0
	entitlementMixinFields1 := entitlementMixin[1].Fields()
	_ = entitlementMixinFields1
	entitlementMixinFields3 := entitlementMixin[3].Fields()
	_ = entitlementMixinFields3
	entitlementFields := schema.Entitlement{}.Fields()
	_ = entitlementFields
	// entitlementDescNamespace is the schema descriptor for namespace field.
	entitlementDescNamespace := entitlementMixinFields1[0].Descriptor()
	// entitlement.NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	entitlement.NamespaceValidator = entitlementDescNamespace.Validators[0].(func(string) error)
	// entitlementDescCreatedAt is the schema descriptor for created_at field.
	entitlementDescCreatedAt := entitlementMixinFields3[0].Descriptor()
	// entitlement.DefaultCreatedAt holds the default value on creation for the created_at field.
	entitlement.DefaultCreatedAt = entitlementDescCreatedAt.Default.(func() time.Time)
	// entitlementDescUpdatedAt is the schema descriptor for updated_at field.
	entitlementDescUpdatedAt := entitlementMixinFields3[1].Descriptor()
	// entitlement.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	entitlement.DefaultUpdatedAt = entitlementDescUpdatedAt.Default.(func() time.Time)
	// entitlement.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	entitlement.UpdateDefaultUpdatedAt = entitlementDescUpdatedAt.UpdateDefault.(func() time.Time)
	// entitlementDescFeatureKey is the schema descriptor for feature_key field.
	entitlementDescFeatureKey := entitlementFields[2].Descriptor()
	// entitlement.FeatureKeyValidator is a validator for the "feature_key" field. It is called by the builders before save.
	entitlement.FeatureKeyValidator = entitlementDescFeatureKey.Validators[0].(func(string) error)
	// entitlementDescSubjectKey is the schema descriptor for subject_key field.
	entitlementDescSubjectKey := entitlementFields[3].Descriptor()
	// entitlement.SubjectKeyValidator is a validator for the "subject_key" field. It is called by the builders before save.
	entitlement.SubjectKeyValidator = entitlementDescSubjectKey.Validators[0].(func(string) error)
	// entitlementDescID is the schema descriptor for id field.
	entitlementDescID := entitlementMixinFields0[0].Descriptor()
	// entitlement.DefaultID holds the default value on creation for the id field.
	entitlement.DefaultID = entitlementDescID.Default.(func() string)
	featureMixin := schema.Feature{}.Mixin()
	featureMixinFields0 := featureMixin[0].Fields()
	_ = featureMixinFields0
	featureMixinFields1 := featureMixin[1].Fields()
	_ = featureMixinFields1
	featureFields := schema.Feature{}.Fields()
	_ = featureFields
	// featureDescCreatedAt is the schema descriptor for created_at field.
	featureDescCreatedAt := featureMixinFields1[0].Descriptor()
	// feature.DefaultCreatedAt holds the default value on creation for the created_at field.
	feature.DefaultCreatedAt = featureDescCreatedAt.Default.(func() time.Time)
	// featureDescUpdatedAt is the schema descriptor for updated_at field.
	featureDescUpdatedAt := featureMixinFields1[1].Descriptor()
	// feature.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feature.DefaultUpdatedAt = featureDescUpdatedAt.Default.(func() time.Time)
	// feature.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feature.UpdateDefaultUpdatedAt = featureDescUpdatedAt.UpdateDefault.(func() time.Time)
	// featureDescNamespace is the schema descriptor for namespace field.
	featureDescNamespace := featureFields[0].Descriptor()
	// feature.NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	feature.NamespaceValidator = featureDescNamespace.Validators[0].(func(string) error)
	// featureDescName is the schema descriptor for name field.
	featureDescName := featureFields[1].Descriptor()
	// feature.NameValidator is a validator for the "name" field. It is called by the builders before save.
	feature.NameValidator = featureDescName.Validators[0].(func(string) error)
	// featureDescKey is the schema descriptor for key field.
	featureDescKey := featureFields[2].Descriptor()
	// feature.KeyValidator is a validator for the "key" field. It is called by the builders before save.
	feature.KeyValidator = featureDescKey.Validators[0].(func(string) error)
	// featureDescID is the schema descriptor for id field.
	featureDescID := featureMixinFields0[0].Descriptor()
	// feature.DefaultID holds the default value on creation for the id field.
	feature.DefaultID = featureDescID.Default.(func() string)
	grantMixin := schema.Grant{}.Mixin()
	grantMixinFields0 := grantMixin[0].Fields()
	_ = grantMixinFields0
	grantMixinFields1 := grantMixin[1].Fields()
	_ = grantMixinFields1
	grantMixinFields3 := grantMixin[3].Fields()
	_ = grantMixinFields3
	grantFields := schema.Grant{}.Fields()
	_ = grantFields
	// grantDescNamespace is the schema descriptor for namespace field.
	grantDescNamespace := grantMixinFields1[0].Descriptor()
	// grant.NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	grant.NamespaceValidator = grantDescNamespace.Validators[0].(func(string) error)
	// grantDescCreatedAt is the schema descriptor for created_at field.
	grantDescCreatedAt := grantMixinFields3[0].Descriptor()
	// grant.DefaultCreatedAt holds the default value on creation for the created_at field.
	grant.DefaultCreatedAt = grantDescCreatedAt.Default.(func() time.Time)
	// grantDescUpdatedAt is the schema descriptor for updated_at field.
	grantDescUpdatedAt := grantMixinFields3[1].Descriptor()
	// grant.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	grant.DefaultUpdatedAt = grantDescUpdatedAt.Default.(func() time.Time)
	// grant.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	grant.UpdateDefaultUpdatedAt = grantDescUpdatedAt.UpdateDefault.(func() time.Time)
	// grantDescPriority is the schema descriptor for priority field.
	grantDescPriority := grantFields[2].Descriptor()
	// grant.DefaultPriority holds the default value on creation for the priority field.
	grant.DefaultPriority = grantDescPriority.Default.(uint8)
	// grantDescID is the schema descriptor for id field.
	grantDescID := grantMixinFields0[0].Descriptor()
	// grant.DefaultID holds the default value on creation for the id field.
	grant.DefaultID = grantDescID.Default.(func() string)
	usageresetMixin := schema.UsageReset{}.Mixin()
	usageresetMixinFields0 := usageresetMixin[0].Fields()
	_ = usageresetMixinFields0
	usageresetMixinFields1 := usageresetMixin[1].Fields()
	_ = usageresetMixinFields1
	usageresetMixinFields2 := usageresetMixin[2].Fields()
	_ = usageresetMixinFields2
	usageresetFields := schema.UsageReset{}.Fields()
	_ = usageresetFields
	// usageresetDescNamespace is the schema descriptor for namespace field.
	usageresetDescNamespace := usageresetMixinFields1[0].Descriptor()
	// usagereset.NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	usagereset.NamespaceValidator = usageresetDescNamespace.Validators[0].(func(string) error)
	// usageresetDescCreatedAt is the schema descriptor for created_at field.
	usageresetDescCreatedAt := usageresetMixinFields2[0].Descriptor()
	// usagereset.DefaultCreatedAt holds the default value on creation for the created_at field.
	usagereset.DefaultCreatedAt = usageresetDescCreatedAt.Default.(func() time.Time)
	// usageresetDescUpdatedAt is the schema descriptor for updated_at field.
	usageresetDescUpdatedAt := usageresetMixinFields2[1].Descriptor()
	// usagereset.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	usagereset.DefaultUpdatedAt = usageresetDescUpdatedAt.Default.(func() time.Time)
	// usagereset.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	usagereset.UpdateDefaultUpdatedAt = usageresetDescUpdatedAt.UpdateDefault.(func() time.Time)
	// usageresetDescID is the schema descriptor for id field.
	usageresetDescID := usageresetMixinFields0[0].Descriptor()
	// usagereset.DefaultID holds the default value on creation for the id field.
	usagereset.DefaultID = usageresetDescID.Default.(func() string)
}
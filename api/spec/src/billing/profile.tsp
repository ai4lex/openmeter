using TypeSpec.Http;
using TypeSpec.OpenAPI;

namespace OpenMeter.Billing;

@route("/api/v1/billing/profile")
@tag("Billing (Experimental)")
interface Profiles {
  /**
   * List all billing profiles
   */
  @get
  @operationId("billingListProfiles")
  list(
    @query
    includeArchived?: boolean = false,

    ...OpenMeter.QueryPagination,
    ...OpenMeter.QueryLimitOffset,
    ...OpenMeter.QueryOrdering<ProfileOrderBy>,
  ): OpenMeter.PaginatedResponse<Profile> | OpenMeter.CommonErrors;

  /**
   * Create a new billing profile
   *
   * Billing profiles are representations of a customer's billing information. Customer overrides
   * can be applied to a billing profile to customize the billing behavior for a specific customer.
   */
  @post
  @summary("Create a new billing profile")
  @operationId("billingCreateProfile")
  create(@body profile: ProfileCreateInput): {
    @statusCode _: 201;
    @body profile: Profile;
  } | OpenMeter.CommonErrors;

  /**
   * Archive a billing profile
   */
  @route("/{id}")
  @delete
  @summary("Archive a billing profile")
  @operationId("billingArchiveProfile")
  archive(@path id: string): {
    @statusCode _: 204;
  } | OpenMeter.CommonErrors | OpenMeter.NotFoundError;

  /**
   * Get a billing profile by ID
   */
  @route("/{id}")
  @get
  @summary("Get a billing profile by ID")
  @operationId("billingGetProfile")
  get(
    @path id: string,
  ): Profile | OpenMeter.NotFoundError | OpenMeter.CommonErrors;

  /**
   * Update a billing profile
   */
  @route("/{id}")
  @put
  @summary("Update a billing profile")
  @operationId("billingUpdateProfile")
  update(@path id: string, @body profile: Profile): {
    @statusCode _: 200;
    @body profile: Profile;
  } | OpenMeter.NotFoundError | OpenMeter.CommonErrors;

  // Customer overrides
}

/**
 * ProfileOrderBy specifies the ordering options for profiles
 */
@friendlyName("BillingProfileOrderBy")
enum ProfileOrderBy {
  createdAt: "createdAt",
  updatedAt: "updatedAt",
}

/**
 * Profile represents a billing profile
 */
@friendlyName("BillingProfile")
model Profile {
  ...global.Resource;

  @summary("The name and contact information for the supplier this billing profile represents")
  supplier: Invoices.Party;

  @summary("The billing workflow settings for this profile")
  workflow: Workflow;

  @summary("Is this the default profile?")
  default: boolean;
}

/**
 * Workflow represents a billing workflow
 */
@friendlyName("BillingWorkflow")
model Workflow {
  ...ResourceTimestamps;
  ...WorkflowSettings;

  /**
   * A unique identifier for the resource.
   */
  @visibility("read")
  @example("01G65Z755AFWAKHE12NY0CQ9FH")
  @summary("ID")
  id: ULID;
}

@friendlyName("BillingWorkflowSettings")
model WorkflowSettings {
  /**
   * When to collect the pending line items into an invoice.
   */
  collectionAlignment?: CollectionAlignment = CollectionAlignment.subscription;

  /**
   * The period for collecting the pending line items into an invoice.
   */
  @encode(DurationKnownEncoding.ISO8601)
  @example("P1D")
  itemCollectionPeriod?: string = "PT1H";

  invoice?: WorkflowInvoice;

  @summary("The tax app used for this workflow")
  taxApp: OpenMeter.App.App;

  @summary("The invoicing app used for this workflow")
  invoicingApp: OpenMeter.App.App;

  @summary("The payment app used for this workflow")
  paymentApp: OpenMeter.App.App;
}

/**
 * AppReference can be used to reference an app during creation only.
 */
@friendlyName("BillingWorkflowAppReference")
model AppReference {
  @summary("The ID of the app, if not specified the type is used to find the default app")
  id?: ULID;

  @summary("The type of the app, if specified the default app is used")
  type?: OpenMeter.App.Type;
}

/**
 * ProfileCreateInput represents the input for creating a billing profile
 */
@friendlyName("BillingProfileCreateInput")
model ProfileCreateInput {
  ...OmitProperties<Profile, "workflow">;
  workflow: WorkflowCreateInput;
}

/**
 * WorkflowCreate represents the workflow settings for creation of a billing profile
 */
@friendlyName("BillingWorkflowCreateInput")
model WorkflowCreateInput {
  ...OmitProperties<Workflow, "taxApp" | "invoicingApp" | "paymentApp">;

  // provider settings for creation (update is not supported)
  @visibility("create")
  taxApp?: AppReference;

  @visibility("create")
  invoicingApp?: AppReference;

  @visibility("create")
  paymentApp?: AppReference;
}

/**
 * CollectionAlignment specifies when the pending line items should be collected into
 * an invoice.
 */
@friendlyName("BillingWorkflowCollectionAlignment")
@summary("Collection alignment")
enum CollectionAlignment {
  /**
   * Align the collection to the start of the subscription period.
   */
  subscription: "subscription",
}

/**
 * WorkflowInvoice represents the invoice settings for a billing workflow
 */
@summary("Workflow invoice settings")
@friendlyName("BillingWorkflowInvoice")
model WorkflowInvoice {
  /**
   * Whether to automatically issue the invoice after the draftPeriod has passed.
   */
  autoAdvance?: boolean = true;

  /**
   * The period for the invoice to be kept in draft status for manual reviews.
   */
  @encode(DurationKnownEncoding.ISO8601)
  @example("P1D")
  draftPeriod?: string = "P1D";

  /**
   * The period after which the invoice is due.
   */
  @encode(DurationKnownEncoding.ISO8601)
  @example("P1D")
  dueAfter?: string = "P7D";

  /**
   * The method to collect the invoice.
   */
  collectionMethod?: CollectionMethod = CollectionMethod.chargeAutomatically;

  /**
   * The resolution of the line items in the invoice.
   */
  itemResolution?: ItemResolution = ItemResolution.period;

  /**
   * Whether to create one line item per subject in the invoice.
   */
  itemPerSubject?: boolean = true;
}

/**
 * CollectionMethod specifies how the invoice should be collected (automatic vs manual)
 */
@friendlyName("BillingWorkflowCollectionMethod")
@summary("Collection method")
enum CollectionMethod {
  chargeAutomatically: "charge_automatically",
  sendInvoice: "send_invoice",
}

/**
 * ItemResolution specifies how the line items should be resolved in the invoice
 */
@friendlyName("BillingWorkflowItemResolution")
@summary("Item resolution")
enum ItemResolution {
  /**
   * Create line items for each day in the billing period.
   */
  day: "day",

  /**
   * Create one line item for the entire billing period.
   */
  period: "period",
}

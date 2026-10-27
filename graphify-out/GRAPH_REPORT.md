# Graph Report - Splitz  (2026-10-27)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 1155 nodes · 3271 edges · 52 communities (48 shown, 4 thin omitted)
- Extraction: 97% EXTRACTED · 3% INFERRED · 0% AMBIGUOUS · INFERRED: 87 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `9fce2bd4`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- el
- modern-screenshot.umd.js
- live-browser.js
- initPageChat
- setLiveState
- handleManualEditActivity
- startVariantObserver
- github.com/google/uuid.UUID
- period_service.go
- args
- PeriodsPage
- documentRefForElement
- args
- connectSSE
- renderDesignVisual
- net/http.ResponseWriter
- NewRouter
- api_test.go
- package.json
- createLiveBrowserDomHelpers
- showToast
- onAnnotDown
- captureElementToBlob
- cleanup
- resolveLiveInjectionAnchor
- createLiveBrowserSessionState
- compilerOptions
- App.tsx
- steps
- middleware.go
- with
- mountSvelteComponentVariant
- compilerOptions
- AuthProvider.tsx
- initGlobalBar
- showBar
- BudgetPage.tsx
- scheduleAcceptCleanup
- createInsertPlaceholder
- syncEditBadgeHitProxies
- docker-publish.yml
- build-and-push
- initDesignPanel
- steps
- config.go
- live-browser-ignores.js
- True
- enableInlineEdit
- tsconfig.json
- splitz

## God Nodes (most connected - your core abstractions)
1. `connectSSE()` - 34 edges
2. `resumeSession()` - 33 edges
3. `setLiveState()` - 33 edges
4. `showToast()` - 31 edges
5. `initGlobalBar()` - 30 edges
6. `el()` - 29 edges
7. `buildInsertConfigureRow()` - 28 edges
8. `cleanup()` - 27 edges
9. `handleKeyDown()` - 27 edges
10. `injectSvelteComponentsFromManifest()` - 26 edges

## Surprising Connections (you probably didn't know these)
- `main()` --calls--> `Load()`  [EXTRACTED]
  cmd/api/main.go → internal/config/config.go
- `TestRequireFirebaseUserAddsValidatedIdentityToContext()` --calls--> `IdentityFromContext()`  [INFERRED]
  internal/auth/middleware_test.go → internal/auth/middleware.go
- `main()` --calls--> `NewFirebaseAuth()`  [EXTRACTED]
  cmd/api/main.go → internal/auth/firebase.go
- `main()` --calls--> `NewPostgresPool()`  [EXTRACTED]
  cmd/api/main.go → internal/database/postgres.go
- `main()` --calls--> `NewRouter()`  [EXTRACTED]
  cmd/api/main.go → internal/server/router.go

## Import Cycles
- None detected.

## Communities (52 total, 4 thin omitted)

### Community 0 - "el"
Cohesion: 0.06
Nodes (67): actionLabel(), agentHasWorkInFlight(), agentStatusText(), applyConfigureBarChrome(), barPaletteForTheme(), bindConfigureCountPillTooltip(), bindConfigureInlineControlHover(), bindConfigureModifierPillHover() (+59 more)

### Community 1 - "modern-screenshot.umd.js"
Cohesion: 0.09
Nodes (55): ae(), be(), bt(), Ce(), s(), Ct(), de(), dt() (+47 more)

### Community 2 - "live-browser.js"
Cohesion: 0.06
Nodes (48): applyGlobalBarLabelState(), applyParamDefaults(), applyParamValue(), applySavedParamValues(), bufferToBase64(), buildColorModels(), buildTypographyModels(), collectFontCssText() (+40 more)

### Community 3 - "initPageChat"
Cohesion: 0.08
Nodes (51): armPageChatForTyping(), attachSteerFocusDebug(), attachSteerFocusGuard(), clearSteerAwaitTimer(), clearSteerFocusRecoverTimer(), collapsePageChat(), configureVoiceContext(), expandPageChat() (+43 more)

### Community 4 - "setLiveState"
Cohesion: 0.12
Nodes (46): actOnAgentTarget(), beginNewLiveConfiguration(), cancelEditing(), cancelEditingToPicking(), cancelInsertConfigure(), clearAnnotations(), clearInsertPicking(), cursorForInsertAxis() (+38 more)

### Community 5 - "handleManualEditActivity"
Cohesion: 0.11
Nodes (40): agentTargetBusyReason(), agentTargetOverlayGone(), agentTargetTaken(), claimAgentTarget(), claimAndActOnAgentTarget(), clearStoredManualApplyState(), declineAgentTargetBusy(), declineAgentTargetUnresolvable() (+32 more)

### Community 6 - "startVariantObserver"
Cohesion: 0.11
Nodes (40): buildCyclingRow(), closedClipPath(), closeTunePopover(), commitAcceptedVariantToDom(), completeParameterPublication(), completeSourceInjection(), cycleVariant(), cyclingCounterText() (+32 more)

### Community 7 - "github.com/google/uuid.UUID"
Cohesion: 0.12
Nodes (19): Category, CreateCategoryRequest, UpdateCategoryRequest, CreateMethodRequest, Method, CreateFinancialSummaryRequest, CreateFixedExpenseRequest, CreatePeriodRequest (+11 more)

### Community 8 - "period_service.go"
Cohesion: 0.24
Nodes (4): AddSalaryRequest, AddSalaryResponse, BudgetRequest, SummaryPreviewRequest

### Community 9 - "args"
Cohesion: 0.06
Nodes (33): VITE_API_URL, VITE_FIREBASE_API_KEY, VITE_FIREBASE_APP_ID, VITE_FIREBASE_AUTH_DOMAIN, VITE_FIREBASE_MESSAGING_SENDER_ID, VITE_FIREBASE_PROJECT_ID, VITE_FIREBASE_STORAGE_BUCKET, build (+25 more)

### Community 10 - "PeriodsPage"
Cohesion: 0.16
Nodes (27): createCategory(), createFixedExpense(), createPeriod(), listCategories(), listFixedExpenses(), listHistoricalFixedExpenses(), listMethods(), previewFinancialSummary() (+19 more)

### Community 11 - "documentRefForElement"
Cohesion: 0.08
Nodes (32): addManualContextText(), canRestoreManualEditElement(), collectManualContextPieces(), walk(), contextElementForManualEdit(), copyEditContainerContext(), copyEditLeafContext(), cssIdent() (+24 more)

### Community 12 - "args"
Cohesion: 0.07
Nodes (30): VITE_API_URL, VITE_FIREBASE_API_KEY, VITE_FIREBASE_APP_ID, VITE_FIREBASE_AUTH_DOMAIN, VITE_FIREBASE_MESSAGING_SENDER_ID, VITE_FIREBASE_PROJECT_ID, VITE_FIREBASE_STORAGE_BUCKET, build (+22 more)

### Community 13 - "connectSSE"
Cohesion: 0.14
Nodes (31): applySavedSessionMeta(), clampVariantIndex(), completeParameterGenerationIfReady(), connectSSE(), enterRecoveryWaitingForAnchor(), findActiveSessionSummary(), findAdoptableServerSession(), findAnyVariantsWrapper() (+23 more)

### Community 14 - "renderDesignVisual"
Cohesion: 0.10
Nodes (30): buildCollapsible(), buildListHtml(), buildRadiiModels(), copyToClipboard(), cssSafe(), designEmptyMessage(), escapeHtml(), fontStack() (+22 more)

### Community 15 - "net/http.ResponseWriter"
Cohesion: 0.22
Nodes (11): CategoryHandler, PeriodHandler, Identity, IdentityFromContext(), CalculateBudget(), CalculateCustomBudget(), HealthCheck(), writeJSON() (+3 more)

### Community 16 - "NewRouter"
Cohesion: 0.13
Nodes (19): main(), MethodHandler, NewFirebaseAuth(), NewPostgresPool(), Me(), NewCategoryHandler(), NewMethodHandler(), NewPeriodHandler() (+11 more)

### Community 17 - "api_test.go"
Cohesion: 0.15
Nodes (20): BudgetResponse, CustomBudgetRequest, CustomBudgetResponse, TestRoundUsesTwoDecimalCommercialRounding(), TestNewRouterDoesNotExposeFinancialRoutesWithoutDependencies(), CalculateBudgetByMethod(), CalculateCustomBudget(), assertFloatApprox() (+12 more)

### Community 18 - "package.json"
Cohesion: 0.10
Nodes (24): dependencies, firebase@^12.19.0, react@^19.1.1, react-dom@^19.1.1, devDependencies, @types/react@^19.1.10, @types/react-dom@^19.1.7, typescript@^5.9.2 (+16 more)

### Community 19 - "createLiveBrowserDomHelpers"
Cohesion: 0.13
Nodes (19): createLiveBrowserDomHelpers(), activeElementDeep(), cloneWithoutChrome(), copyTextViaSelection(), cssId(), liveUiRoot(), makeFrozenAnchor(), own() (+11 more)

### Community 20 - "showToast"
Cohesion: 0.15
Nodes (22): abandonForeignSession(), abandonSupersededGo(), applyEditing(), buildLocatorForLeaf(), discardOrphanedSession(), forbiddenManualTextChars(), handleAccept(), handleDiscard() (+14 more)

### Community 21 - "onAnnotDown"
Cohesion: 0.14
Nodes (22): applyPlaceholderDimensions(), beginEditPin(), buildAnnotationsForCapture(), buildPinElement(), cancelEditingPin(), clampPlaceholderSize(), finalizeEditingPin(), initAnnotOverlay() (+14 more)

### Community 22 - "captureElementToBlob"
Cohesion: 0.12
Nodes (20): averageRgb01(), captureChromeNodes(), captureElementFromRenderedAncestor(), captureElementToBlob(), compileShader(), cssColorToRgb01(), dominantRgb01(), findBackdropAncestor() (+12 more)

### Community 23 - "cleanup"
Cohesion: 0.18
Nodes (20): abortSvelteComponentInjection(), cleanup(), cleanupAcceptedSession(), clearHandled(), clearMountErrorCard(), clearScrollY(), clearSession(), hideShaderOverlay() (+12 more)

### Community 24 - "resolveLiveInjectionAnchor"
Cohesion: 0.15
Nodes (20): buildSvelteExpressionTextMap(), buildSveltePropValuesFromLiveElement(), buildSveltePropValuesV2(), cloneWithoutElements(), collectTextNodes(), collectVisibleTexts(), cssEscapeIdent(), elementMatchesOriginalMarkup() (+12 more)

### Community 25 - "createLiveBrowserSessionState"
Cohesion: 0.19
Nodes (17): createLiveBrowserSessionState(), clearHandled(), clearScrollY(), clearSession(), isHandled(), loadSession(), markHandled(), nextCheckpointRevision() (+9 more)

### Community 26 - "compilerOptions"
Cohesion: 0.11
Nodes (18): compilerOptions, allowJs, allowSyntheticDefaultImports, esModuleInterop, forceConsistentCasingInFileNames, isolatedModules, jsx, lib (+10 more)

### Community 27 - "App.tsx"
Cohesion: 0.23
Nodes (11): App(), AppRoute, AuthenticatedApp(), AppShell(), AppShellProps, AuthProvider(), useAuth(), AuthPage() (+3 more)

### Community 28 - "steps"
Cohesion: 0.12
Nodes (16): runs-on, steps, env, AZURE_MONITOR_CONNECTION_STRING, jobs, build-and-test, name, branches (+8 more)

### Community 29 - "middleware.go"
Cohesion: 0.22
Nodes (9): identityContextKey, TokenVerifier, tokenVerifierFunc, claimString(), RequireFirebaseUser(), TestRequireFirebaseUserAddsValidatedIdentityToContext(), TestRequireFirebaseUserRejectsInvalidToken(), TestRequireFirebaseUserRejectsMissingBearerToken() (+1 more)

### Community 30 - "with"
Cohesion: 0.12
Nodes (16): with, build-args, cache, cache-dependency-path, cache-from, cache-to, context, file (+8 more)

### Community 31 - "mountSvelteComponentVariant"
Cohesion: 0.19
Nodes (15): applyOriginalAttrsToSvelteAnchor(), commitAcceptedSvelteComponentToDom(), componentModuleCandidates(), describeMountFailure(), detectDevServerBase(), findInsertAnchorInDom(), findLiveElementForSvelteManifest(), importFirstReachable() (+7 more)

### Community 32 - "compilerOptions"
Cohesion: 0.14
Nodes (13): compilerOptions, allowImportingTsExtensions, lib, module, moduleDetection, moduleResolution, noEmit, skipLibCheck (+5 more)

### Community 33 - "AuthProvider.tsx"
Cohesion: 0.26
Nodes (9): AuthContextValue, Profile, AuthContext, run(), readableError(), firebaseAuth, firebaseConfig, googleProvider (+1 more)

### Community 34 - "initGlobalBar"
Cohesion: 0.24
Nodes (12): applyLiveBarPreference(), hideAgentPollTooltip(), initGlobalBar(), loadDetectScript(), onDetectMessage(), positionPendingDock(), requestDetectScan(), setLiveBarHidden() (+4 more)

### Community 35 - "showBar"
Cohesion: 0.22
Nodes (13): buildInsertPlaceholderSnapshotFromDom(), buildPickedAnchorSnapshot(), captureAndEmit(), checkpointPayload(), dismissToast(), extractContext(), handleGo(), handleInsertCreate() (+5 more)

### Community 36 - "BudgetPage.tsx"
Cohesion: 0.33
Nodes (8): calculateBudget(), BudgetMethod, BudgetRequest, BudgetResponse, BudgetPage(), handleSubmit(), categories, methods

### Community 37 - "scheduleAcceptCleanup"
Cohesion: 0.31
Nodes (11): acceptedDomAlreadyClean(), clearHandledWrapperReloadStamp(), deferredRecoverySuperseded(), ensureAcceptedDomClean(), findAcceptedRuntimeWrappers(), handledWrapperReloadKey(), reloadAfterMissingAcceptedDom(), restoreAcceptedDomFromSnapshot() (+3 more)

### Community 38 - "createInsertPlaceholder"
Cohesion: 0.18
Nodes (11): applyPlaceholderSizingStyles(), buildPlaceholderResizeHandles(), createInsertPlaceholder(), cursorForPlaceholderEdge(), detectInsertAxis(), detectInsertAxisFromStyle(), finalizeInsertSession(), placeholderSizing() (+3 more)

### Community 40 - "syncEditBadgeHitProxies"
Cohesion: 0.27
Nodes (10): bindEditBadgeProxy(), editBadgeProxyTargets(), initEditBadge(), initEditBadgeHitProxies(), positionEditBadge(), proxyMouseEvent(), setImportantStyle(), styleEditBadgeProxy() (+2 more)

### Community 41 - "docker-publish.yml"
Cohesion: 0.20
Nodes (9): env, REGISTRY_NAMESPACE, jobs, test, name, permissions, contents, name (+1 more)

### Community 42 - "build-and-push"
Cohesion: 0.20
Nodes (10): name, needs, runs-on, strategy, context, file, image, build-and-push (+2 more)

### Community 43 - "initDesignPanel"
Cohesion: 0.39
Nodes (8): buildDesignHeader(), designPanelCss(), fetchDesignSystem(), initDesignPanel(), loadDesignPrefs(), renderDesignChrome(), saveDesignPrefs(), toggleDesignPanel()

### Community 44 - "steps"
Cohesion: 0.29
Nodes (8): steps, id, if, name, run, uses, working-directory, steps

### Community 45 - "config.go"
Cohesion: 0.38
Nodes (3): Config, getEnv(), Load()

### Community 46 - "live-browser-ignores.js"
Cohesion: 0.52
Nodes (6): globToRegex(), matchesScope(), normalizeIgnoreRule(), normalizeIgnoreValue(), pageCandidates(), resolveDetectIgnores()

### Community 47 - "True"
Cohesion: 0.33
Nodes (6): branches, branches, True, pull_request, push, workflow_dispatch

### Community 48 - "enableInlineEdit"
Cohesion: 0.40
Nodes (5): collectEditableTextRows(), visit(), enableInlineEdit(), onInlineInput(), wrapMixedContentTextNodes()

## Knowledge Gaps
- **1 isolated node(s):** `splitz`
  These have ≤1 connection - possible missing edges. (Counts symbols only; 169 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `stopFocusOutIntoChrome()` connect `createLiveBrowserDomHelpers` to `initGlobalBar`, `setLiveState`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Are the 7 inferred relationships involving `initGlobalBar()` (e.g. with `hideAgentPollTooltip()` and `onDetectMessage()`) actually correct?**
  _`initGlobalBar()` has 7 INFERRED edges - model-reasoned connections that need verification._
- **What connects `splitz` to the rest of the system?**
  _1 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `el` be split into smaller, more focused modules?**
  _Cohesion score 0.05750658472344162 - nodes in this community are weakly interconnected._
- **Why does `init()` connect `setLiveState` to `el`, `live-browser.js`, `initPageChat`, `initGlobalBar`, `handleManualEditActivity`, `syncEditBadgeHitProxies`, `initDesignPanel`, `connectSSE`, `createLiveBrowserDomHelpers`, `onAnnotDown`?**
  _High betweenness centrality (0.005) - this node is a cross-community bridge._
- **Should `modern-screenshot.umd.js` be split into smaller, more focused modules?**
  _Cohesion score 0.08897243107769423 - nodes in this community are weakly interconnected._
- **Should `live-browser.js` be split into smaller, more focused modules?**
  _Cohesion score 0.06386066763425254 - nodes in this community are weakly interconnected._
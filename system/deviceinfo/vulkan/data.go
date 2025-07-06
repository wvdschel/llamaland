package vulkan

type VulkanDeviceInfo struct {
	Schema       string             `json:"$schema"`
	Capabilities DeviceCapabilities `json:"capabilities"`
	Profiles     Profiles           `json:"profiles"`
}

type Extensions map[string]int

type VkPhysicalDeviceFeatures struct {
	RobustBufferAccess                      bool `json:"robustBufferAccess"`
	FullDrawIndexUint32                     bool `json:"fullDrawIndexUint32"`
	ImageCubeArray                          bool `json:"imageCubeArray"`
	IndependentBlend                        bool `json:"independentBlend"`
	GeometryShader                          bool `json:"geometryShader"`
	TessellationShader                      bool `json:"tessellationShader"`
	SampleRateShading                       bool `json:"sampleRateShading"`
	DualSrcBlend                            bool `json:"dualSrcBlend"`
	LogicOp                                 bool `json:"logicOp"`
	MultiDrawIndirect                       bool `json:"multiDrawIndirect"`
	DrawIndirectFirstInstance               bool `json:"drawIndirectFirstInstance"`
	DepthClamp                              bool `json:"depthClamp"`
	DepthBiasClamp                          bool `json:"depthBiasClamp"`
	FillModeNonSolid                        bool `json:"fillModeNonSolid"`
	DepthBounds                             bool `json:"depthBounds"`
	WideLines                               bool `json:"wideLines"`
	LargePoints                             bool `json:"largePoints"`
	AlphaToOne                              bool `json:"alphaToOne"`
	MultiViewport                           bool `json:"multiViewport"`
	SamplerAnisotropy                       bool `json:"samplerAnisotropy"`
	TextureCompressionETC2                  bool `json:"textureCompressionETC2"`
	TextureCompressionASTCLDR               bool `json:"textureCompressionASTC_LDR"`
	TextureCompressionBC                    bool `json:"textureCompressionBC"`
	OcclusionQueryPrecise                   bool `json:"occlusionQueryPrecise"`
	PipelineStatisticsQuery                 bool `json:"pipelineStatisticsQuery"`
	VertexPipelineStoresAndAtomics          bool `json:"vertexPipelineStoresAndAtomics"`
	FragmentStoresAndAtomics                bool `json:"fragmentStoresAndAtomics"`
	ShaderTessellationAndGeometryPointSize  bool `json:"shaderTessellationAndGeometryPointSize"`
	ShaderImageGatherExtended               bool `json:"shaderImageGatherExtended"`
	ShaderStorageImageExtendedFormats       bool `json:"shaderStorageImageExtendedFormats"`
	ShaderStorageImageMultisample           bool `json:"shaderStorageImageMultisample"`
	ShaderStorageImageReadWithoutFormat     bool `json:"shaderStorageImageReadWithoutFormat"`
	ShaderStorageImageWriteWithoutFormat    bool `json:"shaderStorageImageWriteWithoutFormat"`
	ShaderUniformBufferArrayDynamicIndexing bool `json:"shaderUniformBufferArrayDynamicIndexing"`
	ShaderSampledImageArrayDynamicIndexing  bool `json:"shaderSampledImageArrayDynamicIndexing"`
	ShaderStorageBufferArrayDynamicIndexing bool `json:"shaderStorageBufferArrayDynamicIndexing"`
	ShaderStorageImageArrayDynamicIndexing  bool `json:"shaderStorageImageArrayDynamicIndexing"`
	ShaderClipDistance                      bool `json:"shaderClipDistance"`
	ShaderCullDistance                      bool `json:"shaderCullDistance"`
	ShaderFloat64                           bool `json:"shaderFloat64"`
	ShaderInt64                             bool `json:"shaderInt64"`
	ShaderInt16                             bool `json:"shaderInt16"`
	ShaderResourceResidency                 bool `json:"shaderResourceResidency"`
	ShaderResourceMinLod                    bool `json:"shaderResourceMinLod"`
	SparseBinding                           bool `json:"sparseBinding"`
	SparseResidencyBuffer                   bool `json:"sparseResidencyBuffer"`
	SparseResidencyImage2D                  bool `json:"sparseResidencyImage2D"`
	SparseResidencyImage3D                  bool `json:"sparseResidencyImage3D"`
	SparseResidency2Samples                 bool `json:"sparseResidency2Samples"`
	SparseResidency4Samples                 bool `json:"sparseResidency4Samples"`
	SparseResidency8Samples                 bool `json:"sparseResidency8Samples"`
	SparseResidency16Samples                bool `json:"sparseResidency16Samples"`
	SparseResidencyAliased                  bool `json:"sparseResidencyAliased"`
	VariableMultisampleRate                 bool `json:"variableMultisampleRate"`
	InheritedQueries                        bool `json:"inheritedQueries"`
}

type VkPhysicalDevice4444FormatsFeaturesEXT struct {
	FormatA4R4G4B4 bool `json:"formatA4R4G4B4"`
	FormatA4B4G4R4 bool `json:"formatA4B4G4R4"`
}

type VkPhysicalDeviceAccelerationStructureFeaturesKHR struct {
	AccelerationStructure                                 bool `json:"accelerationStructure"`
	AccelerationStructureCaptureReplay                    bool `json:"accelerationStructureCaptureReplay"`
	AccelerationStructureIndirectBuild                    bool `json:"accelerationStructureIndirectBuild"`
	AccelerationStructureHostCommands                     bool `json:"accelerationStructureHostCommands"`
	DescriptorBindingAccelerationStructureUpdateAfterBind bool `json:"descriptorBindingAccelerationStructureUpdateAfterBind"`
}

type VkPhysicalDeviceAddressBindingReportFeaturesEXT struct {
	ReportAddressBinding bool `json:"reportAddressBinding"`
}

type VkPhysicalDeviceAttachmentFeedbackLoopDynamicStateFeaturesEXT struct {
	AttachmentFeedbackLoopDynamicState bool `json:"attachmentFeedbackLoopDynamicState"`
}

type VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT struct {
	AttachmentFeedbackLoopLayout bool `json:"attachmentFeedbackLoopLayout"`
}

type VkPhysicalDeviceBorderColorSwizzleFeaturesEXT struct {
	BorderColorSwizzle          bool `json:"borderColorSwizzle"`
	BorderColorSwizzleFromImage bool `json:"borderColorSwizzleFromImage"`
}

type VkPhysicalDeviceBufferDeviceAddressFeaturesEXT struct {
	BufferDeviceAddress              bool `json:"bufferDeviceAddress"`
	BufferDeviceAddressCaptureReplay bool `json:"bufferDeviceAddressCaptureReplay"`
	BufferDeviceAddressMultiDevice   bool `json:"bufferDeviceAddressMultiDevice"`
}

type VkPhysicalDeviceColorWriteEnableFeaturesEXT struct {
	ColorWriteEnable bool `json:"colorWriteEnable"`
}

type VkPhysicalDeviceComputeShaderDerivativesFeaturesKHR struct {
	ComputeDerivativeGroupQuads  bool `json:"computeDerivativeGroupQuads"`
	ComputeDerivativeGroupLinear bool `json:"computeDerivativeGroupLinear"`
}

type VkPhysicalDeviceConditionalRenderingFeaturesEXT struct {
	ConditionalRendering          bool `json:"conditionalRendering"`
	InheritedConditionalRendering bool `json:"inheritedConditionalRendering"`
}

type VkPhysicalDeviceCooperativeMatrixFeaturesKHR struct {
	CooperativeMatrix                   bool `json:"cooperativeMatrix"`
	CooperativeMatrixRobustBufferAccess bool `json:"cooperativeMatrixRobustBufferAccess"`
}

type VkPhysicalDeviceCustomBorderColorFeaturesEXT struct {
	CustomBorderColors             bool `json:"customBorderColors"`
	CustomBorderColorWithoutFormat bool `json:"customBorderColorWithoutFormat"`
}

type VkPhysicalDeviceDepthBiasControlFeaturesEXT struct {
	DepthBiasControl                                bool `json:"depthBiasControl"`
	LeastRepresentableValueForceUnormRepresentation bool `json:"leastRepresentableValueForceUnormRepresentation"`
	FloatRepresentation                             bool `json:"floatRepresentation"`
	DepthBiasExact                                  bool `json:"depthBiasExact"`
}

type VkPhysicalDeviceDepthClampControlFeaturesEXT struct {
	DepthClampControl bool `json:"depthClampControl"`
}

type VkPhysicalDeviceDepthClampZeroOneFeaturesKHR struct {
	DepthClampZeroOne bool `json:"depthClampZeroOne"`
}

type VkPhysicalDeviceDepthClipControlFeaturesEXT struct {
	DepthClipControl bool `json:"depthClipControl"`
}

type VkPhysicalDeviceDepthClipEnableFeaturesEXT struct {
	DepthClipEnable bool `json:"depthClipEnable"`
}

type VkPhysicalDeviceDescriptorBufferFeaturesEXT struct {
	DescriptorBuffer                   bool `json:"descriptorBuffer"`
	DescriptorBufferCaptureReplay      bool `json:"descriptorBufferCaptureReplay"`
	DescriptorBufferImageLayoutIgnored bool `json:"descriptorBufferImageLayoutIgnored"`
	DescriptorBufferPushDescriptors    bool `json:"descriptorBufferPushDescriptors"`
}

type VkPhysicalDeviceDeviceGeneratedCommandsFeaturesEXT struct {
	DeviceGeneratedCommands        bool `json:"deviceGeneratedCommands"`
	DynamicGeneratedPipelineLayout bool `json:"dynamicGeneratedPipelineLayout"`
}

type VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT struct {
	DynamicRenderingUnusedAttachments bool `json:"dynamicRenderingUnusedAttachments"`
}

type VkPhysicalDeviceExtendedDynamicState2FeaturesEXT struct {
	ExtendedDynamicState2                   bool `json:"extendedDynamicState2"`
	ExtendedDynamicState2LogicOp            bool `json:"extendedDynamicState2LogicOp"`
	ExtendedDynamicState2PatchControlPoints bool `json:"extendedDynamicState2PatchControlPoints"`
}

type VkPhysicalDeviceExtendedDynamicState3FeaturesEXT struct {
	ExtendedDynamicState3TessellationDomainOrigin         bool `json:"extendedDynamicState3TessellationDomainOrigin"`
	ExtendedDynamicState3DepthClampEnable                 bool `json:"extendedDynamicState3DepthClampEnable"`
	ExtendedDynamicState3PolygonMode                      bool `json:"extendedDynamicState3PolygonMode"`
	ExtendedDynamicState3RasterizationSamples             bool `json:"extendedDynamicState3RasterizationSamples"`
	ExtendedDynamicState3SampleMask                       bool `json:"extendedDynamicState3SampleMask"`
	ExtendedDynamicState3AlphaToCoverageEnable            bool `json:"extendedDynamicState3AlphaToCoverageEnable"`
	ExtendedDynamicState3AlphaToOneEnable                 bool `json:"extendedDynamicState3AlphaToOneEnable"`
	ExtendedDynamicState3LogicOpEnable                    bool `json:"extendedDynamicState3LogicOpEnable"`
	ExtendedDynamicState3ColorBlendEnable                 bool `json:"extendedDynamicState3ColorBlendEnable"`
	ExtendedDynamicState3ColorBlendEquation               bool `json:"extendedDynamicState3ColorBlendEquation"`
	ExtendedDynamicState3ColorWriteMask                   bool `json:"extendedDynamicState3ColorWriteMask"`
	ExtendedDynamicState3RasterizationStream              bool `json:"extendedDynamicState3RasterizationStream"`
	ExtendedDynamicState3ConservativeRasterizationMode    bool `json:"extendedDynamicState3ConservativeRasterizationMode"`
	ExtendedDynamicState3ExtraPrimitiveOverestimationSize bool `json:"extendedDynamicState3ExtraPrimitiveOverestimationSize"`
	ExtendedDynamicState3DepthClipEnable                  bool `json:"extendedDynamicState3DepthClipEnable"`
	ExtendedDynamicState3SampleLocationsEnable            bool `json:"extendedDynamicState3SampleLocationsEnable"`
	ExtendedDynamicState3ColorBlendAdvanced               bool `json:"extendedDynamicState3ColorBlendAdvanced"`
	ExtendedDynamicState3ProvokingVertexMode              bool `json:"extendedDynamicState3ProvokingVertexMode"`
	ExtendedDynamicState3LineRasterizationMode            bool `json:"extendedDynamicState3LineRasterizationMode"`
	ExtendedDynamicState3LineStippleEnable                bool `json:"extendedDynamicState3LineStippleEnable"`
	ExtendedDynamicState3DepthClipNegativeOneToOne        bool `json:"extendedDynamicState3DepthClipNegativeOneToOne"`
	ExtendedDynamicState3ViewportWScalingEnable           bool `json:"extendedDynamicState3ViewportWScalingEnable"`
	ExtendedDynamicState3ViewportSwizzle                  bool `json:"extendedDynamicState3ViewportSwizzle"`
	ExtendedDynamicState3CoverageToColorEnable            bool `json:"extendedDynamicState3CoverageToColorEnable"`
	ExtendedDynamicState3CoverageToColorLocation          bool `json:"extendedDynamicState3CoverageToColorLocation"`
	ExtendedDynamicState3CoverageModulationMode           bool `json:"extendedDynamicState3CoverageModulationMode"`
	ExtendedDynamicState3CoverageModulationTableEnable    bool `json:"extendedDynamicState3CoverageModulationTableEnable"`
	ExtendedDynamicState3CoverageModulationTable          bool `json:"extendedDynamicState3CoverageModulationTable"`
	ExtendedDynamicState3CoverageReductionMode            bool `json:"extendedDynamicState3CoverageReductionMode"`
	ExtendedDynamicState3RepresentativeFragmentTestEnable bool `json:"extendedDynamicState3RepresentativeFragmentTestEnable"`
	ExtendedDynamicState3ShadingRateImageEnable           bool `json:"extendedDynamicState3ShadingRateImageEnable"`
}

type VkPhysicalDeviceExtendedDynamicStateFeaturesEXT struct {
	ExtendedDynamicState bool `json:"extendedDynamicState"`
}

type VkPhysicalDeviceFaultFeaturesEXT struct {
	DeviceFault             bool `json:"deviceFault"`
	DeviceFaultVendorBinary bool `json:"deviceFaultVendorBinary"`
}

type VkPhysicalDeviceFragmentShaderBarycentricFeaturesKHR struct {
	FragmentShaderBarycentric bool `json:"fragmentShaderBarycentric"`
}

type VkPhysicalDeviceFragmentShaderInterlockFeaturesEXT struct {
	FragmentShaderSampleInterlock      bool `json:"fragmentShaderSampleInterlock"`
	FragmentShaderPixelInterlock       bool `json:"fragmentShaderPixelInterlock"`
	FragmentShaderShadingRateInterlock bool `json:"fragmentShaderShadingRateInterlock"`
}

type VkPhysicalDeviceFragmentShadingRateFeaturesKHR struct {
	PipelineFragmentShadingRate   bool `json:"pipelineFragmentShadingRate"`
	PrimitiveFragmentShadingRate  bool `json:"primitiveFragmentShadingRate"`
	AttachmentFragmentShadingRate bool `json:"attachmentFragmentShadingRate"`
}

type VkPhysicalDeviceGraphicsPipelineLibraryFeaturesEXT struct {
	GraphicsPipelineLibrary bool `json:"graphicsPipelineLibrary"`
}

type VkPhysicalDeviceImage2DViewOf3DFeaturesEXT struct {
	Image2DViewOf3D   bool `json:"image2DViewOf3D"`
	Sampler2DViewOf3D bool `json:"sampler2DViewOf3D"`
}

type VkPhysicalDeviceImageCompressionControlFeaturesEXT struct {
	ImageCompressionControl bool `json:"imageCompressionControl"`
}

type VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT struct {
	ImageSlicedViewOf3D bool `json:"imageSlicedViewOf3D"`
}

type VkPhysicalDeviceImageViewMinLodFeaturesEXT struct {
	MinLod bool `json:"minLod"`
}

type VkPhysicalDeviceLegacyVertexAttributesFeaturesEXT struct {
	LegacyVertexAttributes bool `json:"legacyVertexAttributes"`
}

type VkPhysicalDeviceMaintenance7FeaturesKHR struct {
	Maintenance7 bool `json:"maintenance7"`
}

type VkPhysicalDeviceMaintenance8FeaturesKHR struct {
	Maintenance8 bool `json:"maintenance8"`
}

type VkPhysicalDeviceMapMemoryPlacedFeaturesEXT struct {
	MemoryMapPlaced      bool `json:"memoryMapPlaced"`
	MemoryMapRangePlaced bool `json:"memoryMapRangePlaced"`
	MemoryUnmapReserve   bool `json:"memoryUnmapReserve"`
}

type VkPhysicalDeviceMemoryPriorityFeaturesEXT struct {
	MemoryPriority bool `json:"memoryPriority"`
}

type VkPhysicalDeviceMeshShaderFeaturesEXT struct {
	TaskShader                             bool `json:"taskShader"`
	MeshShader                             bool `json:"meshShader"`
	MultiviewMeshShader                    bool `json:"multiviewMeshShader"`
	PrimitiveFragmentShadingRateMeshShader bool `json:"primitiveFragmentShadingRateMeshShader"`
	MeshShaderQueries                      bool `json:"meshShaderQueries"`
}

type VkPhysicalDeviceMultiDrawFeaturesEXT struct {
	MultiDraw bool `json:"multiDraw"`
}

type VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT struct {
	MutableDescriptorType bool `json:"mutableDescriptorType"`
}

type VkPhysicalDeviceNestedCommandBufferFeaturesEXT struct {
	NestedCommandBuffer                bool `json:"nestedCommandBuffer"`
	NestedCommandBufferRendering       bool `json:"nestedCommandBufferRendering"`
	NestedCommandBufferSimultaneousUse bool `json:"nestedCommandBufferSimultaneousUse"`
}

type VkPhysicalDeviceNonSeamlessCubeMapFeaturesEXT struct {
	NonSeamlessCubeMap bool `json:"nonSeamlessCubeMap"`
}

type VkPhysicalDevicePipelineBinaryFeaturesKHR struct {
	PipelineBinaries bool `json:"pipelineBinaries"`
}

type VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR struct {
	PipelineExecutableInfo bool `json:"pipelineExecutableInfo"`
}

type VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT struct {
	PipelineLibraryGroupHandles bool `json:"pipelineLibraryGroupHandles"`
}

type VkPhysicalDevicePresentIDFeaturesKHR struct {
	PresentID bool `json:"presentId"`
}

type VkPhysicalDevicePresentWaitFeaturesKHR struct {
	PresentWait bool `json:"presentWait"`
}

type VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT struct {
	PrimitiveTopologyListRestart      bool `json:"primitiveTopologyListRestart"`
	PrimitiveTopologyPatchListRestart bool `json:"primitiveTopologyPatchListRestart"`
}

type VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT struct {
	PrimitivesGeneratedQuery                      bool `json:"primitivesGeneratedQuery"`
	PrimitivesGeneratedQueryWithRasterizerDiscard bool `json:"primitivesGeneratedQueryWithRasterizerDiscard"`
	PrimitivesGeneratedQueryWithNonZeroStreams    bool `json:"primitivesGeneratedQueryWithNonZeroStreams"`
}

type VkPhysicalDeviceProvokingVertexFeaturesEXT struct {
	ProvokingVertexLast                       bool `json:"provokingVertexLast"`
	TransformFeedbackPreservesProvokingVertex bool `json:"transformFeedbackPreservesProvokingVertex"`
}

type VkPhysicalDeviceRayQueryFeaturesKHR struct {
	RayQuery bool `json:"rayQuery"`
}

type VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR struct {
	RayTracingMaintenance1               bool `json:"rayTracingMaintenance1"`
	RayTracingPipelineTraceRaysIndirect2 bool `json:"rayTracingPipelineTraceRaysIndirect2"`
}

type VkPhysicalDeviceRayTracingPipelineFeaturesKHR struct {
	RayTracingPipeline                                    bool `json:"rayTracingPipeline"`
	RayTracingPipelineShaderGroupHandleCaptureReplay      bool `json:"rayTracingPipelineShaderGroupHandleCaptureReplay"`
	RayTracingPipelineShaderGroupHandleCaptureReplayMixed bool `json:"rayTracingPipelineShaderGroupHandleCaptureReplayMixed"`
	RayTracingPipelineTraceRaysIndirect                   bool `json:"rayTracingPipelineTraceRaysIndirect"`
	RayTraversalPrimitiveCulling                          bool `json:"rayTraversalPrimitiveCulling"`
}

type VkPhysicalDeviceRayTracingPositionFetchFeaturesKHR struct {
	RayTracingPositionFetch bool `json:"rayTracingPositionFetch"`
}

type VkPhysicalDeviceRobustness2FeaturesEXT struct {
	RobustBufferAccess2 bool `json:"robustBufferAccess2"`
	RobustImageAccess2  bool `json:"robustImageAccess2"`
	NullDescriptor      bool `json:"nullDescriptor"`
}

type VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT struct {
	ShaderBufferFloat16Atomics      bool `json:"shaderBufferFloat16Atomics"`
	ShaderBufferFloat16AtomicAdd    bool `json:"shaderBufferFloat16AtomicAdd"`
	ShaderBufferFloat16AtomicMinMax bool `json:"shaderBufferFloat16AtomicMinMax"`
	ShaderBufferFloat32AtomicMinMax bool `json:"shaderBufferFloat32AtomicMinMax"`
	ShaderBufferFloat64AtomicMinMax bool `json:"shaderBufferFloat64AtomicMinMax"`
	ShaderSharedFloat16Atomics      bool `json:"shaderSharedFloat16Atomics"`
	ShaderSharedFloat16AtomicAdd    bool `json:"shaderSharedFloat16AtomicAdd"`
	ShaderSharedFloat16AtomicMinMax bool `json:"shaderSharedFloat16AtomicMinMax"`
	ShaderSharedFloat32AtomicMinMax bool `json:"shaderSharedFloat32AtomicMinMax"`
	ShaderSharedFloat64AtomicMinMax bool `json:"shaderSharedFloat64AtomicMinMax"`
	ShaderImageFloat32AtomicMinMax  bool `json:"shaderImageFloat32AtomicMinMax"`
	SparseImageFloat32AtomicMinMax  bool `json:"sparseImageFloat32AtomicMinMax"`
}

type VkPhysicalDeviceShaderAtomicFloatFeaturesEXT struct {
	ShaderBufferFloat32Atomics   bool `json:"shaderBufferFloat32Atomics"`
	ShaderBufferFloat32AtomicAdd bool `json:"shaderBufferFloat32AtomicAdd"`
	ShaderBufferFloat64Atomics   bool `json:"shaderBufferFloat64Atomics"`
	ShaderBufferFloat64AtomicAdd bool `json:"shaderBufferFloat64AtomicAdd"`
	ShaderSharedFloat32Atomics   bool `json:"shaderSharedFloat32Atomics"`
	ShaderSharedFloat32AtomicAdd bool `json:"shaderSharedFloat32AtomicAdd"`
	ShaderSharedFloat64Atomics   bool `json:"shaderSharedFloat64Atomics"`
	ShaderSharedFloat64AtomicAdd bool `json:"shaderSharedFloat64AtomicAdd"`
	ShaderImageFloat32Atomics    bool `json:"shaderImageFloat32Atomics"`
	ShaderImageFloat32AtomicAdd  bool `json:"shaderImageFloat32AtomicAdd"`
	SparseImageFloat32Atomics    bool `json:"sparseImageFloat32Atomics"`
	SparseImageFloat32AtomicAdd  bool `json:"sparseImageFloat32AtomicAdd"`
}

type VkPhysicalDeviceShaderClockFeaturesKHR struct {
	ShaderSubgroupClock bool `json:"shaderSubgroupClock"`
	ShaderDeviceClock   bool `json:"shaderDeviceClock"`
}

type VkPhysicalDeviceShaderImageAtomicInt64FeaturesEXT struct {
	ShaderImageInt64Atomics bool `json:"shaderImageInt64Atomics"`
	SparseImageInt64Atomics bool `json:"sparseImageInt64Atomics"`
}

type VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR struct {
	ShaderMaximalReconvergence bool `json:"shaderMaximalReconvergence"`
}

type VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT struct {
	ShaderModuleIdentifier bool `json:"shaderModuleIdentifier"`
}

type VkPhysicalDeviceShaderObjectFeaturesEXT struct {
	ShaderObject bool `json:"shaderObject"`
}

type VkPhysicalDeviceShaderQuadControlFeaturesKHR struct {
	ShaderQuadControl bool `json:"shaderQuadControl"`
}

type VkPhysicalDeviceShaderRelaxedExtendedInstructionFeaturesKHR struct {
	ShaderRelaxedExtendedInstruction bool `json:"shaderRelaxedExtendedInstruction"`
}

type VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT struct {
	ShaderReplicatedComposites bool `json:"shaderReplicatedComposites"`
}

type VkPhysicalDeviceShaderSubgroupUniformControlFlowFeaturesKHR struct {
	ShaderSubgroupUniformControlFlow bool `json:"shaderSubgroupUniformControlFlow"`
}

type VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT struct {
	SwapchainMaintenance1 bool `json:"swapchainMaintenance1"`
}

type VkPhysicalDeviceTexelBufferAlignmentFeaturesEXT struct {
	TexelBufferAlignment bool `json:"texelBufferAlignment"`
}

type VkPhysicalDeviceTransformFeedbackFeaturesEXT struct {
	TransformFeedback bool `json:"transformFeedback"`
	GeometryStreams   bool `json:"geometryStreams"`
}

type VkPhysicalDeviceVertexInputDynamicStateFeaturesEXT struct {
	VertexInputDynamicState bool `json:"vertexInputDynamicState"`
}

type VkPhysicalDeviceVideoMaintenance1FeaturesKHR struct {
	VideoMaintenance1 bool `json:"videoMaintenance1"`
}

type VkPhysicalDeviceVulkan11Features struct {
	StorageBuffer16BitAccess           bool `json:"storageBuffer16BitAccess"`
	UniformAndStorageBuffer16BitAccess bool `json:"uniformAndStorageBuffer16BitAccess"`
	StoragePushConstant16              bool `json:"storagePushConstant16"`
	StorageInputOutput16               bool `json:"storageInputOutput16"`
	Multiview                          bool `json:"multiview"`
	MultiviewGeometryShader            bool `json:"multiviewGeometryShader"`
	MultiviewTessellationShader        bool `json:"multiviewTessellationShader"`
	VariablePointersStorageBuffer      bool `json:"variablePointersStorageBuffer"`
	VariablePointers                   bool `json:"variablePointers"`
	ProtectedMemory                    bool `json:"protectedMemory"`
	SamplerYcbcrConversion             bool `json:"samplerYcbcrConversion"`
	ShaderDrawParameters               bool `json:"shaderDrawParameters"`
}

type VkPhysicalDeviceVulkan12Features struct {
	SamplerMirrorClampToEdge                           bool `json:"samplerMirrorClampToEdge"`
	DrawIndirectCount                                  bool `json:"drawIndirectCount"`
	StorageBuffer8BitAccess                            bool `json:"storageBuffer8BitAccess"`
	UniformAndStorageBuffer8BitAccess                  bool `json:"uniformAndStorageBuffer8BitAccess"`
	StoragePushConstant8                               bool `json:"storagePushConstant8"`
	ShaderBufferInt64Atomics                           bool `json:"shaderBufferInt64Atomics"`
	ShaderSharedInt64Atomics                           bool `json:"shaderSharedInt64Atomics"`
	ShaderFloat16                                      bool `json:"shaderFloat16"`
	ShaderInt8                                         bool `json:"shaderInt8"`
	DescriptorIndexing                                 bool `json:"descriptorIndexing"`
	ShaderInputAttachmentArrayDynamicIndexing          bool `json:"shaderInputAttachmentArrayDynamicIndexing"`
	ShaderUniformTexelBufferArrayDynamicIndexing       bool `json:"shaderUniformTexelBufferArrayDynamicIndexing"`
	ShaderStorageTexelBufferArrayDynamicIndexing       bool `json:"shaderStorageTexelBufferArrayDynamicIndexing"`
	ShaderUniformBufferArrayNonUniformIndexing         bool `json:"shaderUniformBufferArrayNonUniformIndexing"`
	ShaderSampledImageArrayNonUniformIndexing          bool `json:"shaderSampledImageArrayNonUniformIndexing"`
	ShaderStorageBufferArrayNonUniformIndexing         bool `json:"shaderStorageBufferArrayNonUniformIndexing"`
	ShaderStorageImageArrayNonUniformIndexing          bool `json:"shaderStorageImageArrayNonUniformIndexing"`
	ShaderInputAttachmentArrayNonUniformIndexing       bool `json:"shaderInputAttachmentArrayNonUniformIndexing"`
	ShaderUniformTexelBufferArrayNonUniformIndexing    bool `json:"shaderUniformTexelBufferArrayNonUniformIndexing"`
	ShaderStorageTexelBufferArrayNonUniformIndexing    bool `json:"shaderStorageTexelBufferArrayNonUniformIndexing"`
	DescriptorBindingUniformBufferUpdateAfterBind      bool `json:"descriptorBindingUniformBufferUpdateAfterBind"`
	DescriptorBindingSampledImageUpdateAfterBind       bool `json:"descriptorBindingSampledImageUpdateAfterBind"`
	DescriptorBindingStorageImageUpdateAfterBind       bool `json:"descriptorBindingStorageImageUpdateAfterBind"`
	DescriptorBindingStorageBufferUpdateAfterBind      bool `json:"descriptorBindingStorageBufferUpdateAfterBind"`
	DescriptorBindingUniformTexelBufferUpdateAfterBind bool `json:"descriptorBindingUniformTexelBufferUpdateAfterBind"`
	DescriptorBindingStorageTexelBufferUpdateAfterBind bool `json:"descriptorBindingStorageTexelBufferUpdateAfterBind"`
	DescriptorBindingUpdateUnusedWhilePending          bool `json:"descriptorBindingUpdateUnusedWhilePending"`
	DescriptorBindingPartiallyBound                    bool `json:"descriptorBindingPartiallyBound"`
	DescriptorBindingVariableDescriptorCount           bool `json:"descriptorBindingVariableDescriptorCount"`
	RuntimeDescriptorArray                             bool `json:"runtimeDescriptorArray"`
	SamplerFilterMinmax                                bool `json:"samplerFilterMinmax"`
	ScalarBlockLayout                                  bool `json:"scalarBlockLayout"`
	ImagelessFramebuffer                               bool `json:"imagelessFramebuffer"`
	UniformBufferStandardLayout                        bool `json:"uniformBufferStandardLayout"`
	ShaderSubgroupExtendedTypes                        bool `json:"shaderSubgroupExtendedTypes"`
	SeparateDepthStencilLayouts                        bool `json:"separateDepthStencilLayouts"`
	HostQueryReset                                     bool `json:"hostQueryReset"`
	TimelineSemaphore                                  bool `json:"timelineSemaphore"`
	BufferDeviceAddress                                bool `json:"bufferDeviceAddress"`
	BufferDeviceAddressCaptureReplay                   bool `json:"bufferDeviceAddressCaptureReplay"`
	BufferDeviceAddressMultiDevice                     bool `json:"bufferDeviceAddressMultiDevice"`
	VulkanMemoryModel                                  bool `json:"vulkanMemoryModel"`
	VulkanMemoryModelDeviceScope                       bool `json:"vulkanMemoryModelDeviceScope"`
	VulkanMemoryModelAvailabilityVisibilityChains      bool `json:"vulkanMemoryModelAvailabilityVisibilityChains"`
	ShaderOutputViewportIndex                          bool `json:"shaderOutputViewportIndex"`
	ShaderOutputLayer                                  bool `json:"shaderOutputLayer"`
	SubgroupBroadcastDynamicID                         bool `json:"subgroupBroadcastDynamicId"`
}

type VkPhysicalDeviceVulkan13Features struct {
	RobustImageAccess                                  bool `json:"robustImageAccess"`
	InlineUniformBlock                                 bool `json:"inlineUniformBlock"`
	DescriptorBindingInlineUniformBlockUpdateAfterBind bool `json:"descriptorBindingInlineUniformBlockUpdateAfterBind"`
	PipelineCreationCacheControl                       bool `json:"pipelineCreationCacheControl"`
	PrivateData                                        bool `json:"privateData"`
	ShaderDemoteToHelperInvocation                     bool `json:"shaderDemoteToHelperInvocation"`
	ShaderTerminateInvocation                          bool `json:"shaderTerminateInvocation"`
	SubgroupSizeControl                                bool `json:"subgroupSizeControl"`
	ComputeFullSubgroups                               bool `json:"computeFullSubgroups"`
	Synchronization2                                   bool `json:"synchronization2"`
	TextureCompressionASTCHDR                          bool `json:"textureCompressionASTC_HDR"`
	ShaderZeroInitializeWorkgroupMemory                bool `json:"shaderZeroInitializeWorkgroupMemory"`
	DynamicRendering                                   bool `json:"dynamicRendering"`
	ShaderIntegerDotProduct                            bool `json:"shaderIntegerDotProduct"`
	Maintenance4                                       bool `json:"maintenance4"`
}

type VkPhysicalDeviceVulkan14Features struct {
	GlobalPriorityQuery                    bool `json:"globalPriorityQuery"`
	ShaderSubgroupRotate                   bool `json:"shaderSubgroupRotate"`
	ShaderSubgroupRotateClustered          bool `json:"shaderSubgroupRotateClustered"`
	ShaderFloatControls2                   bool `json:"shaderFloatControls2"`
	ShaderExpectAssume                     bool `json:"shaderExpectAssume"`
	RectangularLines                       bool `json:"rectangularLines"`
	BresenhamLines                         bool `json:"bresenhamLines"`
	SmoothLines                            bool `json:"smoothLines"`
	StippledRectangularLines               bool `json:"stippledRectangularLines"`
	StippledBresenhamLines                 bool `json:"stippledBresenhamLines"`
	StippledSmoothLines                    bool `json:"stippledSmoothLines"`
	VertexAttributeInstanceRateDivisor     bool `json:"vertexAttributeInstanceRateDivisor"`
	VertexAttributeInstanceRateZeroDivisor bool `json:"vertexAttributeInstanceRateZeroDivisor"`
	IndexTypeUint8                         bool `json:"indexTypeUint8"`
	DynamicRenderingLocalRead              bool `json:"dynamicRenderingLocalRead"`
	Maintenance5                           bool `json:"maintenance5"`
	Maintenance6                           bool `json:"maintenance6"`
	PipelineProtectedAccess                bool `json:"pipelineProtectedAccess"`
	PipelineRobustness                     bool `json:"pipelineRobustness"`
	HostImageCopy                          bool `json:"hostImageCopy"`
	PushDescriptor                         bool `json:"pushDescriptor"`
}

type VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR struct {
	WorkgroupMemoryExplicitLayout                  bool `json:"workgroupMemoryExplicitLayout"`
	WorkgroupMemoryExplicitLayoutScalarBlockLayout bool `json:"workgroupMemoryExplicitLayoutScalarBlockLayout"`
	WorkgroupMemoryExplicitLayout8BitAccess        bool `json:"workgroupMemoryExplicitLayout8BitAccess"`
	WorkgroupMemoryExplicitLayout16BitAccess       bool `json:"workgroupMemoryExplicitLayout16BitAccess"`
}

type VkPhysicalDeviceYcbcrImageArraysFeaturesEXT struct {
	YcbcrImageArrays bool `json:"ycbcrImageArrays"`
}

type Features struct {
	VkPhysicalDeviceFeatures                                      VkPhysicalDeviceFeatures                                      `json:"VkPhysicalDeviceFeatures"`
	VkPhysicalDevice4444FormatsFeaturesEXT                        VkPhysicalDevice4444FormatsFeaturesEXT                        `json:"VkPhysicalDevice4444FormatsFeaturesEXT"`
	VkPhysicalDeviceAccelerationStructureFeaturesKHR              VkPhysicalDeviceAccelerationStructureFeaturesKHR              `json:"VkPhysicalDeviceAccelerationStructureFeaturesKHR"`
	VkPhysicalDeviceAddressBindingReportFeaturesEXT               VkPhysicalDeviceAddressBindingReportFeaturesEXT               `json:"VkPhysicalDeviceAddressBindingReportFeaturesEXT"`
	VkPhysicalDeviceAttachmentFeedbackLoopDynamicStateFeaturesEXT VkPhysicalDeviceAttachmentFeedbackLoopDynamicStateFeaturesEXT `json:"VkPhysicalDeviceAttachmentFeedbackLoopDynamicStateFeaturesEXT"`
	VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT       VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT       `json:"VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT"`
	VkPhysicalDeviceBorderColorSwizzleFeaturesEXT                 VkPhysicalDeviceBorderColorSwizzleFeaturesEXT                 `json:"VkPhysicalDeviceBorderColorSwizzleFeaturesEXT"`
	VkPhysicalDeviceBufferDeviceAddressFeaturesEXT                VkPhysicalDeviceBufferDeviceAddressFeaturesEXT                `json:"VkPhysicalDeviceBufferDeviceAddressFeaturesEXT"`
	VkPhysicalDeviceColorWriteEnableFeaturesEXT                   VkPhysicalDeviceColorWriteEnableFeaturesEXT                   `json:"VkPhysicalDeviceColorWriteEnableFeaturesEXT"`
	VkPhysicalDeviceComputeShaderDerivativesFeaturesKHR           VkPhysicalDeviceComputeShaderDerivativesFeaturesKHR           `json:"VkPhysicalDeviceComputeShaderDerivativesFeaturesKHR"`
	VkPhysicalDeviceConditionalRenderingFeaturesEXT               VkPhysicalDeviceConditionalRenderingFeaturesEXT               `json:"VkPhysicalDeviceConditionalRenderingFeaturesEXT"`
	VkPhysicalDeviceCooperativeMatrixFeaturesKHR                  VkPhysicalDeviceCooperativeMatrixFeaturesKHR                  `json:"VkPhysicalDeviceCooperativeMatrixFeaturesKHR"`
	VkPhysicalDeviceCustomBorderColorFeaturesEXT                  VkPhysicalDeviceCustomBorderColorFeaturesEXT                  `json:"VkPhysicalDeviceCustomBorderColorFeaturesEXT"`
	VkPhysicalDeviceDepthBiasControlFeaturesEXT                   VkPhysicalDeviceDepthBiasControlFeaturesEXT                   `json:"VkPhysicalDeviceDepthBiasControlFeaturesEXT"`
	VkPhysicalDeviceDepthClampControlFeaturesEXT                  VkPhysicalDeviceDepthClampControlFeaturesEXT                  `json:"VkPhysicalDeviceDepthClampControlFeaturesEXT"`
	VkPhysicalDeviceDepthClampZeroOneFeaturesKHR                  VkPhysicalDeviceDepthClampZeroOneFeaturesKHR                  `json:"VkPhysicalDeviceDepthClampZeroOneFeaturesKHR"`
	VkPhysicalDeviceDepthClipControlFeaturesEXT                   VkPhysicalDeviceDepthClipControlFeaturesEXT                   `json:"VkPhysicalDeviceDepthClipControlFeaturesEXT"`
	VkPhysicalDeviceDepthClipEnableFeaturesEXT                    VkPhysicalDeviceDepthClipEnableFeaturesEXT                    `json:"VkPhysicalDeviceDepthClipEnableFeaturesEXT"`
	VkPhysicalDeviceDescriptorBufferFeaturesEXT                   VkPhysicalDeviceDescriptorBufferFeaturesEXT                   `json:"VkPhysicalDeviceDescriptorBufferFeaturesEXT"`
	VkPhysicalDeviceDeviceGeneratedCommandsFeaturesEXT            VkPhysicalDeviceDeviceGeneratedCommandsFeaturesEXT            `json:"VkPhysicalDeviceDeviceGeneratedCommandsFeaturesEXT"`
	VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT  VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT  `json:"VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT"`
	VkPhysicalDeviceExtendedDynamicState2FeaturesEXT              VkPhysicalDeviceExtendedDynamicState2FeaturesEXT              `json:"VkPhysicalDeviceExtendedDynamicState2FeaturesEXT"`
	VkPhysicalDeviceExtendedDynamicState3FeaturesEXT              VkPhysicalDeviceExtendedDynamicState3FeaturesEXT              `json:"VkPhysicalDeviceExtendedDynamicState3FeaturesEXT"`
	VkPhysicalDeviceExtendedDynamicStateFeaturesEXT               VkPhysicalDeviceExtendedDynamicStateFeaturesEXT               `json:"VkPhysicalDeviceExtendedDynamicStateFeaturesEXT"`
	VkPhysicalDeviceFaultFeaturesEXT                              VkPhysicalDeviceFaultFeaturesEXT                              `json:"VkPhysicalDeviceFaultFeaturesEXT"`
	VkPhysicalDeviceFragmentShaderBarycentricFeaturesKHR          VkPhysicalDeviceFragmentShaderBarycentricFeaturesKHR          `json:"VkPhysicalDeviceFragmentShaderBarycentricFeaturesKHR"`
	VkPhysicalDeviceFragmentShaderInterlockFeaturesEXT            VkPhysicalDeviceFragmentShaderInterlockFeaturesEXT            `json:"VkPhysicalDeviceFragmentShaderInterlockFeaturesEXT"`
	VkPhysicalDeviceFragmentShadingRateFeaturesKHR                VkPhysicalDeviceFragmentShadingRateFeaturesKHR                `json:"VkPhysicalDeviceFragmentShadingRateFeaturesKHR"`
	VkPhysicalDeviceGraphicsPipelineLibraryFeaturesEXT            VkPhysicalDeviceGraphicsPipelineLibraryFeaturesEXT            `json:"VkPhysicalDeviceGraphicsPipelineLibraryFeaturesEXT"`
	VkPhysicalDeviceImage2DViewOf3DFeaturesEXT                    VkPhysicalDeviceImage2DViewOf3DFeaturesEXT                    `json:"VkPhysicalDeviceImage2DViewOf3DFeaturesEXT"`
	VkPhysicalDeviceImageCompressionControlFeaturesEXT            VkPhysicalDeviceImageCompressionControlFeaturesEXT            `json:"VkPhysicalDeviceImageCompressionControlFeaturesEXT"`
	VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT                VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT                `json:"VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT"`
	VkPhysicalDeviceImageViewMinLodFeaturesEXT                    VkPhysicalDeviceImageViewMinLodFeaturesEXT                    `json:"VkPhysicalDeviceImageViewMinLodFeaturesEXT"`
	VkPhysicalDeviceLegacyVertexAttributesFeaturesEXT             VkPhysicalDeviceLegacyVertexAttributesFeaturesEXT             `json:"VkPhysicalDeviceLegacyVertexAttributesFeaturesEXT"`
	VkPhysicalDeviceMaintenance7FeaturesKHR                       VkPhysicalDeviceMaintenance7FeaturesKHR                       `json:"VkPhysicalDeviceMaintenance7FeaturesKHR"`
	VkPhysicalDeviceMaintenance8FeaturesKHR                       VkPhysicalDeviceMaintenance8FeaturesKHR                       `json:"VkPhysicalDeviceMaintenance8FeaturesKHR"`
	VkPhysicalDeviceMapMemoryPlacedFeaturesEXT                    VkPhysicalDeviceMapMemoryPlacedFeaturesEXT                    `json:"VkPhysicalDeviceMapMemoryPlacedFeaturesEXT"`
	VkPhysicalDeviceMemoryPriorityFeaturesEXT                     VkPhysicalDeviceMemoryPriorityFeaturesEXT                     `json:"VkPhysicalDeviceMemoryPriorityFeaturesEXT"`
	VkPhysicalDeviceMeshShaderFeaturesEXT                         VkPhysicalDeviceMeshShaderFeaturesEXT                         `json:"VkPhysicalDeviceMeshShaderFeaturesEXT"`
	VkPhysicalDeviceMultiDrawFeaturesEXT                          VkPhysicalDeviceMultiDrawFeaturesEXT                          `json:"VkPhysicalDeviceMultiDrawFeaturesEXT"`
	VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT              VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT              `json:"VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT"`
	VkPhysicalDeviceNestedCommandBufferFeaturesEXT                VkPhysicalDeviceNestedCommandBufferFeaturesEXT                `json:"VkPhysicalDeviceNestedCommandBufferFeaturesEXT"`
	VkPhysicalDeviceNonSeamlessCubeMapFeaturesEXT                 VkPhysicalDeviceNonSeamlessCubeMapFeaturesEXT                 `json:"VkPhysicalDeviceNonSeamlessCubeMapFeaturesEXT"`
	VkPhysicalDevicePipelineBinaryFeaturesKHR                     VkPhysicalDevicePipelineBinaryFeaturesKHR                     `json:"VkPhysicalDevicePipelineBinaryFeaturesKHR"`
	VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR       VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR       `json:"VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR"`
	VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT        VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT        `json:"VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT"`
	VkPhysicalDevicePresentIDFeaturesKHR                          VkPhysicalDevicePresentIDFeaturesKHR                          `json:"VkPhysicalDevicePresentIdFeaturesKHR"`
	VkPhysicalDevicePresentWaitFeaturesKHR                        VkPhysicalDevicePresentWaitFeaturesKHR                        `json:"VkPhysicalDevicePresentWaitFeaturesKHR"`
	VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT       VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT       `json:"VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT"`
	VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT           VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT           `json:"VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT"`
	VkPhysicalDeviceProvokingVertexFeaturesEXT                    VkPhysicalDeviceProvokingVertexFeaturesEXT                    `json:"VkPhysicalDeviceProvokingVertexFeaturesEXT"`
	VkPhysicalDeviceRayQueryFeaturesKHR                           VkPhysicalDeviceRayQueryFeaturesKHR                           `json:"VkPhysicalDeviceRayQueryFeaturesKHR"`
	VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR             VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR             `json:"VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR"`
	VkPhysicalDeviceRayTracingPipelineFeaturesKHR                 VkPhysicalDeviceRayTracingPipelineFeaturesKHR                 `json:"VkPhysicalDeviceRayTracingPipelineFeaturesKHR"`
	VkPhysicalDeviceRayTracingPositionFetchFeaturesKHR            VkPhysicalDeviceRayTracingPositionFetchFeaturesKHR            `json:"VkPhysicalDeviceRayTracingPositionFetchFeaturesKHR"`
	VkPhysicalDeviceRobustness2FeaturesEXT                        VkPhysicalDeviceRobustness2FeaturesEXT                        `json:"VkPhysicalDeviceRobustness2FeaturesEXT"`
	VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT                 VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT                 `json:"VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT"`
	VkPhysicalDeviceShaderAtomicFloatFeaturesEXT                  VkPhysicalDeviceShaderAtomicFloatFeaturesEXT                  `json:"VkPhysicalDeviceShaderAtomicFloatFeaturesEXT"`
	VkPhysicalDeviceShaderClockFeaturesKHR                        VkPhysicalDeviceShaderClockFeaturesKHR                        `json:"VkPhysicalDeviceShaderClockFeaturesKHR"`
	VkPhysicalDeviceShaderImageAtomicInt64FeaturesEXT             VkPhysicalDeviceShaderImageAtomicInt64FeaturesEXT             `json:"VkPhysicalDeviceShaderImageAtomicInt64FeaturesEXT"`
	VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR         VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR         `json:"VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR"`
	VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT             VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT             `json:"VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT"`
	VkPhysicalDeviceShaderObjectFeaturesEXT                       VkPhysicalDeviceShaderObjectFeaturesEXT                       `json:"VkPhysicalDeviceShaderObjectFeaturesEXT"`
	VkPhysicalDeviceShaderQuadControlFeaturesKHR                  VkPhysicalDeviceShaderQuadControlFeaturesKHR                  `json:"VkPhysicalDeviceShaderQuadControlFeaturesKHR"`
	VkPhysicalDeviceShaderRelaxedExtendedInstructionFeaturesKHR   VkPhysicalDeviceShaderRelaxedExtendedInstructionFeaturesKHR   `json:"VkPhysicalDeviceShaderRelaxedExtendedInstructionFeaturesKHR"`
	VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT         VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT         `json:"VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT"`
	VkPhysicalDeviceShaderSubgroupUniformControlFlowFeaturesKHR   VkPhysicalDeviceShaderSubgroupUniformControlFlowFeaturesKHR   `json:"VkPhysicalDeviceShaderSubgroupUniformControlFlowFeaturesKHR"`
	VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT              VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT              `json:"VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT"`
	VkPhysicalDeviceTexelBufferAlignmentFeaturesEXT               VkPhysicalDeviceTexelBufferAlignmentFeaturesEXT               `json:"VkPhysicalDeviceTexelBufferAlignmentFeaturesEXT"`
	VkPhysicalDeviceTransformFeedbackFeaturesEXT                  VkPhysicalDeviceTransformFeedbackFeaturesEXT                  `json:"VkPhysicalDeviceTransformFeedbackFeaturesEXT"`
	VkPhysicalDeviceVertexInputDynamicStateFeaturesEXT            VkPhysicalDeviceVertexInputDynamicStateFeaturesEXT            `json:"VkPhysicalDeviceVertexInputDynamicStateFeaturesEXT"`
	VkPhysicalDeviceVideoMaintenance1FeaturesKHR                  VkPhysicalDeviceVideoMaintenance1FeaturesKHR                  `json:"VkPhysicalDeviceVideoMaintenance1FeaturesKHR"`
	VkPhysicalDeviceVulkan11Features                              VkPhysicalDeviceVulkan11Features                              `json:"VkPhysicalDeviceVulkan11Features"`
	VkPhysicalDeviceVulkan12Features                              VkPhysicalDeviceVulkan12Features                              `json:"VkPhysicalDeviceVulkan12Features"`
	VkPhysicalDeviceVulkan13Features                              VkPhysicalDeviceVulkan13Features                              `json:"VkPhysicalDeviceVulkan13Features"`
	VkPhysicalDeviceVulkan14Features                              VkPhysicalDeviceVulkan14Features                              `json:"VkPhysicalDeviceVulkan14Features"`
	VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR      VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR      `json:"VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR"`
	VkPhysicalDeviceYcbcrImageArraysFeaturesEXT                   VkPhysicalDeviceYcbcrImageArraysFeaturesEXT                   `json:"VkPhysicalDeviceYcbcrImageArraysFeaturesEXT"`
}

type Limits struct {
	MaxImageDimension1D                             int       `json:"maxImageDimension1D"`
	MaxImageDimension2D                             int       `json:"maxImageDimension2D"`
	MaxImageDimension3D                             int       `json:"maxImageDimension3D"`
	MaxImageDimensionCube                           int       `json:"maxImageDimensionCube"`
	MaxImageArrayLayers                             int       `json:"maxImageArrayLayers"`
	MaxTexelBufferElements                          int64     `json:"maxTexelBufferElements"`
	MaxUniformBufferRange                           int64     `json:"maxUniformBufferRange"`
	MaxStorageBufferRange                           int64     `json:"maxStorageBufferRange"`
	MaxPushConstantsSize                            int       `json:"maxPushConstantsSize"`
	MaxMemoryAllocationCount                        int64     `json:"maxMemoryAllocationCount"`
	MaxSamplerAllocationCount                       int       `json:"maxSamplerAllocationCount"`
	BufferImageGranularity                          int       `json:"bufferImageGranularity"`
	SparseAddressSpaceSize                          int64     `json:"sparseAddressSpaceSize"`
	MaxBoundDescriptorSets                          int       `json:"maxBoundDescriptorSets"`
	MaxPerStageDescriptorSamplers                   int       `json:"maxPerStageDescriptorSamplers"`
	MaxPerStageDescriptorUniformBuffers             int       `json:"maxPerStageDescriptorUniformBuffers"`
	MaxPerStageDescriptorStorageBuffers             int       `json:"maxPerStageDescriptorStorageBuffers"`
	MaxPerStageDescriptorSampledImages              int       `json:"maxPerStageDescriptorSampledImages"`
	MaxPerStageDescriptorStorageImages              int       `json:"maxPerStageDescriptorStorageImages"`
	MaxPerStageDescriptorInputAttachments           int       `json:"maxPerStageDescriptorInputAttachments"`
	MaxPerStageResources                            int       `json:"maxPerStageResources"`
	MaxDescriptorSetSamplers                        int       `json:"maxDescriptorSetSamplers"`
	MaxDescriptorSetUniformBuffers                  int       `json:"maxDescriptorSetUniformBuffers"`
	MaxDescriptorSetUniformBuffersDynamic           int       `json:"maxDescriptorSetUniformBuffersDynamic"`
	MaxDescriptorSetStorageBuffers                  int       `json:"maxDescriptorSetStorageBuffers"`
	MaxDescriptorSetStorageBuffersDynamic           int       `json:"maxDescriptorSetStorageBuffersDynamic"`
	MaxDescriptorSetSampledImages                   int       `json:"maxDescriptorSetSampledImages"`
	MaxDescriptorSetStorageImages                   int       `json:"maxDescriptorSetStorageImages"`
	MaxDescriptorSetInputAttachments                int       `json:"maxDescriptorSetInputAttachments"`
	MaxVertexInputAttributes                        int       `json:"maxVertexInputAttributes"`
	MaxVertexInputBindings                          int       `json:"maxVertexInputBindings"`
	MaxVertexInputAttributeOffset                   int64     `json:"maxVertexInputAttributeOffset"`
	MaxVertexInputBindingStride                     int       `json:"maxVertexInputBindingStride"`
	MaxVertexOutputComponents                       int       `json:"maxVertexOutputComponents"`
	MaxTessellationGenerationLevel                  int       `json:"maxTessellationGenerationLevel"`
	MaxTessellationPatchSize                        int       `json:"maxTessellationPatchSize"`
	MaxTessellationControlPerVertexInputComponents  int       `json:"maxTessellationControlPerVertexInputComponents"`
	MaxTessellationControlPerVertexOutputComponents int       `json:"maxTessellationControlPerVertexOutputComponents"`
	MaxTessellationControlPerPatchOutputComponents  int       `json:"maxTessellationControlPerPatchOutputComponents"`
	MaxTessellationControlTotalOutputComponents     int       `json:"maxTessellationControlTotalOutputComponents"`
	MaxTessellationEvaluationInputComponents        int       `json:"maxTessellationEvaluationInputComponents"`
	MaxTessellationEvaluationOutputComponents       int       `json:"maxTessellationEvaluationOutputComponents"`
	MaxGeometryShaderInvocations                    int       `json:"maxGeometryShaderInvocations"`
	MaxGeometryInputComponents                      int       `json:"maxGeometryInputComponents"`
	MaxGeometryOutputComponents                     int       `json:"maxGeometryOutputComponents"`
	MaxGeometryOutputVertices                       int       `json:"maxGeometryOutputVertices"`
	MaxGeometryTotalOutputComponents                int       `json:"maxGeometryTotalOutputComponents"`
	MaxFragmentInputComponents                      int       `json:"maxFragmentInputComponents"`
	MaxFragmentOutputAttachments                    int       `json:"maxFragmentOutputAttachments"`
	MaxFragmentDualSrcAttachments                   int       `json:"maxFragmentDualSrcAttachments"`
	MaxFragmentCombinedOutputResources              int       `json:"maxFragmentCombinedOutputResources"`
	MaxComputeSharedMemorySize                      int       `json:"maxComputeSharedMemorySize"`
	MaxComputeWorkGroupCount                        []any     `json:"maxComputeWorkGroupCount"`
	MaxComputeWorkGroupInvocations                  int       `json:"maxComputeWorkGroupInvocations"`
	MaxComputeWorkGroupSize                         []int     `json:"maxComputeWorkGroupSize"`
	SubPixelPrecisionBits                           int       `json:"subPixelPrecisionBits"`
	SubTexelPrecisionBits                           int       `json:"subTexelPrecisionBits"`
	MipmapPrecisionBits                             int       `json:"mipmapPrecisionBits"`
	MaxDrawIndexedIndexValue                        int64     `json:"maxDrawIndexedIndexValue"`
	MaxDrawIndirectCount                            int64     `json:"maxDrawIndirectCount"`
	MaxSamplerLodBias                               int       `json:"maxSamplerLodBias"`
	MaxSamplerAnisotropy                            int       `json:"maxSamplerAnisotropy"`
	MaxViewports                                    int       `json:"maxViewports"`
	MaxViewportDimensions                           []int     `json:"maxViewportDimensions"`
	ViewportBoundsRange                             []int     `json:"viewportBoundsRange"`
	ViewportSubPixelBits                            int       `json:"viewportSubPixelBits"`
	MinMemoryMapAlignment                           int       `json:"minMemoryMapAlignment"`
	MinTexelBufferOffsetAlignment                   int       `json:"minTexelBufferOffsetAlignment"`
	MinUniformBufferOffsetAlignment                 int       `json:"minUniformBufferOffsetAlignment"`
	MinStorageBufferOffsetAlignment                 int       `json:"minStorageBufferOffsetAlignment"`
	MinTexelOffset                                  int       `json:"minTexelOffset"`
	MaxTexelOffset                                  int       `json:"maxTexelOffset"`
	MinTexelGatherOffset                            int       `json:"minTexelGatherOffset"`
	MaxTexelGatherOffset                            int       `json:"maxTexelGatherOffset"`
	MinInterpolationOffset                          int       `json:"minInterpolationOffset"`
	MaxInterpolationOffset                          int       `json:"maxInterpolationOffset"`
	SubPixelInterpolationOffsetBits                 int       `json:"subPixelInterpolationOffsetBits"`
	MaxFramebufferWidth                             int       `json:"maxFramebufferWidth"`
	MaxFramebufferHeight                            int       `json:"maxFramebufferHeight"`
	MaxFramebufferLayers                            int       `json:"maxFramebufferLayers"`
	FramebufferColorSampleCounts                    []string  `json:"framebufferColorSampleCounts"`
	FramebufferDepthSampleCounts                    []string  `json:"framebufferDepthSampleCounts"`
	FramebufferStencilSampleCounts                  []string  `json:"framebufferStencilSampleCounts"`
	FramebufferNoAttachmentsSampleCounts            []string  `json:"framebufferNoAttachmentsSampleCounts"`
	MaxColorAttachments                             int       `json:"maxColorAttachments"`
	SampledImageColorSampleCounts                   []string  `json:"sampledImageColorSampleCounts"`
	SampledImageIntegerSampleCounts                 []string  `json:"sampledImageIntegerSampleCounts"`
	SampledImageDepthSampleCounts                   []string  `json:"sampledImageDepthSampleCounts"`
	SampledImageStencilSampleCounts                 []string  `json:"sampledImageStencilSampleCounts"`
	StorageImageSampleCounts                        []string  `json:"storageImageSampleCounts"`
	MaxSampleMaskWords                              int       `json:"maxSampleMaskWords"`
	TimestampComputeAndGraphics                     bool      `json:"timestampComputeAndGraphics"`
	TimestampPeriod                                 int       `json:"timestampPeriod"`
	MaxClipDistances                                int       `json:"maxClipDistances"`
	MaxCullDistances                                int       `json:"maxCullDistances"`
	MaxCombinedClipAndCullDistances                 int       `json:"maxCombinedClipAndCullDistances"`
	DiscreteQueuePriorities                         int       `json:"discreteQueuePriorities"`
	PointSizeRange                                  []float64 `json:"pointSizeRange"`
	LineWidthRange                                  []int     `json:"lineWidthRange"`
	PointSizeGranularity                            float64   `json:"pointSizeGranularity"`
	LineWidthGranularity                            float64   `json:"lineWidthGranularity"`
	StrictLines                                     bool      `json:"strictLines"`
	StandardSampleLocations                         bool      `json:"standardSampleLocations"`
	OptimalBufferCopyOffsetAlignment                int       `json:"optimalBufferCopyOffsetAlignment"`
	OptimalBufferCopyRowPitchAlignment              int       `json:"optimalBufferCopyRowPitchAlignment"`
	NonCoherentAtomSize                             int       `json:"nonCoherentAtomSize"`
}

type SparseProperties struct {
	ResidencyStandard2DBlockShape            bool `json:"residencyStandard2DBlockShape"`
	ResidencyStandard2DMultisampleBlockShape bool `json:"residencyStandard2DMultisampleBlockShape"`
	ResidencyStandard3DBlockShape            bool `json:"residencyStandard3DBlockShape"`
	ResidencyAlignedMipSize                  bool `json:"residencyAlignedMipSize"`
	ResidencyNonResidentStrict               bool `json:"residencyNonResidentStrict"`
}

type VkPhysicalDeviceProperties struct {
	APIVersion        int              `json:"apiVersion"`
	DeviceID          int              `json:"deviceID"`
	DeviceName        string           `json:"deviceName"`
	DeviceType        string           `json:"deviceType"`
	DriverVersion     int              `json:"driverVersion"`
	Limits            Limits           `json:"limits"`
	PipelineCacheUUID []int            `json:"pipelineCacheUUID"`
	SparseProperties  SparseProperties `json:"sparseProperties"`
	VendorID          int              `json:"vendorID"`
}

type VkPhysicalDeviceAccelerationStructurePropertiesKHR struct {
	MaxGeometryCount                                           int `json:"maxGeometryCount"`
	MaxInstanceCount                                           int `json:"maxInstanceCount"`
	MaxPrimitiveCount                                          int `json:"maxPrimitiveCount"`
	MaxPerStageDescriptorAccelerationStructures                int `json:"maxPerStageDescriptorAccelerationStructures"`
	MaxPerStageDescriptorUpdateAfterBindAccelerationStructures int `json:"maxPerStageDescriptorUpdateAfterBindAccelerationStructures"`
	MaxDescriptorSetAccelerationStructures                     int `json:"maxDescriptorSetAccelerationStructures"`
	MaxDescriptorSetUpdateAfterBindAccelerationStructures      int `json:"maxDescriptorSetUpdateAfterBindAccelerationStructures"`
	MinAccelerationStructureScratchOffsetAlignment             int `json:"minAccelerationStructureScratchOffsetAlignment"`
}

type VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR struct {
	MeshAndTaskShaderDerivatives bool `json:"meshAndTaskShaderDerivatives"`
}

type VkPhysicalDeviceConservativeRasterizationPropertiesEXT struct {
	PrimitiveOverestimationSize                 int  `json:"primitiveOverestimationSize"`
	MaxExtraPrimitiveOverestimationSize         int  `json:"maxExtraPrimitiveOverestimationSize"`
	ExtraPrimitiveOverestimationSizeGranularity int  `json:"extraPrimitiveOverestimationSizeGranularity"`
	PrimitiveUnderestimation                    bool `json:"primitiveUnderestimation"`
	ConservativePointAndLineRasterization       bool `json:"conservativePointAndLineRasterization"`
	DegenerateTrianglesRasterized               bool `json:"degenerateTrianglesRasterized"`
	DegenerateLinesRasterized                   bool `json:"degenerateLinesRasterized"`
	FullyCoveredFragmentShaderInputVariable     bool `json:"fullyCoveredFragmentShaderInputVariable"`
	ConservativeRasterizationPostDepthCoverage  bool `json:"conservativeRasterizationPostDepthCoverage"`
}

type VkPhysicalDeviceCooperativeMatrixPropertiesKHR struct {
	CooperativeMatrixSupportedStages []string `json:"cooperativeMatrixSupportedStages"`
}

type VkPhysicalDeviceCustomBorderColorPropertiesEXT struct {
	MaxCustomBorderColorSamplers int `json:"maxCustomBorderColorSamplers"`
}

type VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT struct {
	CombinedImageSamplerDensityMapDescriptorSize int `json:"combinedImageSamplerDensityMapDescriptorSize"`
}

type VkPhysicalDeviceDescriptorBufferPropertiesEXT struct {
	CombinedImageSamplerDescriptorSingleArray            bool  `json:"combinedImageSamplerDescriptorSingleArray"`
	BufferlessPushDescriptors                            bool  `json:"bufferlessPushDescriptors"`
	AllowSamplerImageViewPostSubmitCreation              bool  `json:"allowSamplerImageViewPostSubmitCreation"`
	DescriptorBufferOffsetAlignment                      int   `json:"descriptorBufferOffsetAlignment"`
	MaxDescriptorBufferBindings                          int   `json:"maxDescriptorBufferBindings"`
	MaxResourceDescriptorBufferBindings                  int   `json:"maxResourceDescriptorBufferBindings"`
	MaxSamplerDescriptorBufferBindings                   int   `json:"maxSamplerDescriptorBufferBindings"`
	MaxEmbeddedImmutableSamplerBindings                  int   `json:"maxEmbeddedImmutableSamplerBindings"`
	MaxEmbeddedImmutableSamplers                         int   `json:"maxEmbeddedImmutableSamplers"`
	BufferCaptureReplayDescriptorDataSize                int   `json:"bufferCaptureReplayDescriptorDataSize"`
	ImageCaptureReplayDescriptorDataSize                 int   `json:"imageCaptureReplayDescriptorDataSize"`
	ImageViewCaptureReplayDescriptorDataSize             int   `json:"imageViewCaptureReplayDescriptorDataSize"`
	SamplerCaptureReplayDescriptorDataSize               int   `json:"samplerCaptureReplayDescriptorDataSize"`
	AccelerationStructureCaptureReplayDescriptorDataSize int   `json:"accelerationStructureCaptureReplayDescriptorDataSize"`
	SamplerDescriptorSize                                int   `json:"samplerDescriptorSize"`
	CombinedImageSamplerDescriptorSize                   int   `json:"combinedImageSamplerDescriptorSize"`
	SampledImageDescriptorSize                           int   `json:"sampledImageDescriptorSize"`
	StorageImageDescriptorSize                           int   `json:"storageImageDescriptorSize"`
	UniformTexelBufferDescriptorSize                     int   `json:"uniformTexelBufferDescriptorSize"`
	RobustUniformTexelBufferDescriptorSize               int   `json:"robustUniformTexelBufferDescriptorSize"`
	StorageTexelBufferDescriptorSize                     int   `json:"storageTexelBufferDescriptorSize"`
	RobustStorageTexelBufferDescriptorSize               int   `json:"robustStorageTexelBufferDescriptorSize"`
	UniformBufferDescriptorSize                          int   `json:"uniformBufferDescriptorSize"`
	RobustUniformBufferDescriptorSize                    int   `json:"robustUniformBufferDescriptorSize"`
	StorageBufferDescriptorSize                          int   `json:"storageBufferDescriptorSize"`
	RobustStorageBufferDescriptorSize                    int   `json:"robustStorageBufferDescriptorSize"`
	InputAttachmentDescriptorSize                        int   `json:"inputAttachmentDescriptorSize"`
	AccelerationStructureDescriptorSize                  int   `json:"accelerationStructureDescriptorSize"`
	MaxSamplerDescriptorBufferRange                      int64 `json:"maxSamplerDescriptorBufferRange"`
	MaxResourceDescriptorBufferRange                     int64 `json:"maxResourceDescriptorBufferRange"`
	SamplerDescriptorBufferAddressSpaceSize              int64 `json:"samplerDescriptorBufferAddressSpaceSize"`
	ResourceDescriptorBufferAddressSpaceSize             int64 `json:"resourceDescriptorBufferAddressSpaceSize"`
	DescriptorBufferAddressSpaceSize                     int64 `json:"descriptorBufferAddressSpaceSize"`
}

type VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT struct {
	MaxIndirectPipelineCount                             int      `json:"maxIndirectPipelineCount"`
	MaxIndirectShaderObjectCount                         int      `json:"maxIndirectShaderObjectCount"`
	MaxIndirectSequenceCount                             int      `json:"maxIndirectSequenceCount"`
	MaxIndirectCommandsTokenCount                        int      `json:"maxIndirectCommandsTokenCount"`
	MaxIndirectCommandsTokenOffset                       int      `json:"maxIndirectCommandsTokenOffset"`
	MaxIndirectCommandsIndirectStride                    int      `json:"maxIndirectCommandsIndirectStride"`
	SupportedIndirectCommandsInputModes                  []string `json:"supportedIndirectCommandsInputModes"`
	SupportedIndirectCommandsShaderStages                []string `json:"supportedIndirectCommandsShaderStages"`
	SupportedIndirectCommandsShaderStagesPipelineBinding []string `json:"supportedIndirectCommandsShaderStagesPipelineBinding"`
	SupportedIndirectCommandsShaderStagesShaderBinding   []string `json:"supportedIndirectCommandsShaderStagesShaderBinding"`
	DeviceGeneratedCommandsTransformFeedback             bool     `json:"deviceGeneratedCommandsTransformFeedback"`
	DeviceGeneratedCommandsMultiDrawIndirectCount        bool     `json:"deviceGeneratedCommandsMultiDrawIndirectCount"`
}

type VkPhysicalDeviceDiscardRectanglePropertiesEXT struct {
	MaxDiscardRectangles int `json:"maxDiscardRectangles"`
}

type VkPhysicalDeviceDrmPropertiesEXT struct {
	HasPrimary   bool `json:"hasPrimary"`
	HasRender    bool `json:"hasRender"`
	PrimaryMajor int  `json:"primaryMajor"`
	PrimaryMinor int  `json:"primaryMinor"`
	RenderMajor  int  `json:"renderMajor"`
	RenderMinor  int  `json:"renderMinor"`
}

type VkPhysicalDeviceExtendedDynamicState3PropertiesEXT struct {
	DynamicPrimitiveTopologyUnrestricted bool `json:"dynamicPrimitiveTopologyUnrestricted"`
}

type VkPhysicalDeviceExternalMemoryHostPropertiesEXT struct {
	MinImportedHostPointerAlignment int `json:"minImportedHostPointerAlignment"`
}

type VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR struct {
	TriStripVertexOrderIndependentOfProvokingVertex bool `json:"triStripVertexOrderIndependentOfProvokingVertex"`
}

type MinFragmentShadingRateAttachmentTexelSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type MaxFragmentShadingRateAttachmentTexelSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type MaxFragmentSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type VkPhysicalDeviceFragmentShadingRatePropertiesKHR struct {
	MinFragmentShadingRateAttachmentTexelSize            MinFragmentShadingRateAttachmentTexelSize `json:"minFragmentShadingRateAttachmentTexelSize"`
	MaxFragmentShadingRateAttachmentTexelSize            MaxFragmentShadingRateAttachmentTexelSize `json:"maxFragmentShadingRateAttachmentTexelSize"`
	MaxFragmentShadingRateAttachmentTexelSizeAspectRatio int                                       `json:"maxFragmentShadingRateAttachmentTexelSizeAspectRatio"`
	PrimitiveFragmentShadingRateWithMultipleViewports    bool                                      `json:"primitiveFragmentShadingRateWithMultipleViewports"`
	LayeredShadingRateAttachments                        bool                                      `json:"layeredShadingRateAttachments"`
	FragmentShadingRateNonTrivialCombinerOps             bool                                      `json:"fragmentShadingRateNonTrivialCombinerOps"`
	MaxFragmentSize                                      MaxFragmentSize                           `json:"maxFragmentSize"`
	MaxFragmentSizeAspectRatio                           int                                       `json:"maxFragmentSizeAspectRatio"`
	MaxFragmentShadingRateCoverageSamples                int                                       `json:"maxFragmentShadingRateCoverageSamples"`
	MaxFragmentShadingRateRasterizationSamples           string                                    `json:"maxFragmentShadingRateRasterizationSamples"`
	FragmentShadingRateWithShaderDepthStencilWrites      bool                                      `json:"fragmentShadingRateWithShaderDepthStencilWrites"`
	FragmentShadingRateWithSampleMask                    bool                                      `json:"fragmentShadingRateWithSampleMask"`
	FragmentShadingRateWithShaderSampleMask              bool                                      `json:"fragmentShadingRateWithShaderSampleMask"`
	FragmentShadingRateWithConservativeRasterization     bool                                      `json:"fragmentShadingRateWithConservativeRasterization"`
	FragmentShadingRateWithFragmentShaderInterlock       bool                                      `json:"fragmentShadingRateWithFragmentShaderInterlock"`
	FragmentShadingRateWithCustomSampleLocations         bool                                      `json:"fragmentShadingRateWithCustomSampleLocations"`
	FragmentShadingRateStrictMultiplyCombiner            bool                                      `json:"fragmentShadingRateStrictMultiplyCombiner"`
}

type VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT struct {
	GraphicsPipelineLibraryFastLinking                        bool `json:"graphicsPipelineLibraryFastLinking"`
	GraphicsPipelineLibraryIndependentInterpolationDecoration bool `json:"graphicsPipelineLibraryIndependentInterpolationDecoration"`
}

type VkPhysicalDeviceLayeredAPIPropertiesListKHR struct {
	LayeredAPICount int    `json:"layeredApiCount"`
	PLayeredApis    string `json:"pLayeredApis"`
}

type VkPhysicalDeviceLegacyVertexAttributesPropertiesEXT struct {
	NativeUnalignedPerformance bool `json:"nativeUnalignedPerformance"`
}

type VkPhysicalDeviceMaintenance7PropertiesKHR struct {
	RobustFragmentShadingRateAttachmentAccess                 bool `json:"robustFragmentShadingRateAttachmentAccess"`
	SeparateDepthStencilAttachmentAccess                      bool `json:"separateDepthStencilAttachmentAccess"`
	MaxDescriptorSetTotalUniformBuffersDynamic                int  `json:"maxDescriptorSetTotalUniformBuffersDynamic"`
	MaxDescriptorSetTotalStorageBuffersDynamic                int  `json:"maxDescriptorSetTotalStorageBuffersDynamic"`
	MaxDescriptorSetTotalBuffersDynamic                       int  `json:"maxDescriptorSetTotalBuffersDynamic"`
	MaxDescriptorSetUpdateAfterBindTotalUniformBuffersDynamic int  `json:"maxDescriptorSetUpdateAfterBindTotalUniformBuffersDynamic"`
	MaxDescriptorSetUpdateAfterBindTotalStorageBuffersDynamic int  `json:"maxDescriptorSetUpdateAfterBindTotalStorageBuffersDynamic"`
	MaxDescriptorSetUpdateAfterBindTotalBuffersDynamic        int  `json:"maxDescriptorSetUpdateAfterBindTotalBuffersDynamic"`
}

type VkPhysicalDeviceMapMemoryPlacedPropertiesEXT struct {
	MinPlacedMemoryMapAlignment int `json:"minPlacedMemoryMapAlignment"`
}

type VkPhysicalDeviceMeshShaderPropertiesEXT struct {
	MaxTaskWorkGroupTotalCount            int   `json:"maxTaskWorkGroupTotalCount"`
	MaxTaskWorkGroupCount                 []int `json:"maxTaskWorkGroupCount"`
	MaxTaskWorkGroupInvocations           int   `json:"maxTaskWorkGroupInvocations"`
	MaxTaskWorkGroupSize                  []int `json:"maxTaskWorkGroupSize"`
	MaxTaskPayloadSize                    int   `json:"maxTaskPayloadSize"`
	MaxTaskSharedMemorySize               int   `json:"maxTaskSharedMemorySize"`
	MaxTaskPayloadAndSharedMemorySize     int   `json:"maxTaskPayloadAndSharedMemorySize"`
	MaxMeshWorkGroupTotalCount            int   `json:"maxMeshWorkGroupTotalCount"`
	MaxMeshWorkGroupCount                 []int `json:"maxMeshWorkGroupCount"`
	MaxMeshWorkGroupInvocations           int   `json:"maxMeshWorkGroupInvocations"`
	MaxMeshWorkGroupSize                  []int `json:"maxMeshWorkGroupSize"`
	MaxMeshSharedMemorySize               int   `json:"maxMeshSharedMemorySize"`
	MaxMeshPayloadAndSharedMemorySize     int   `json:"maxMeshPayloadAndSharedMemorySize"`
	MaxMeshOutputMemorySize               int   `json:"maxMeshOutputMemorySize"`
	MaxMeshPayloadAndOutputMemorySize     int   `json:"maxMeshPayloadAndOutputMemorySize"`
	MaxMeshOutputComponents               int   `json:"maxMeshOutputComponents"`
	MaxMeshOutputVertices                 int   `json:"maxMeshOutputVertices"`
	MaxMeshOutputPrimitives               int   `json:"maxMeshOutputPrimitives"`
	MaxMeshOutputLayers                   int   `json:"maxMeshOutputLayers"`
	MaxMeshMultiviewViewCount             int   `json:"maxMeshMultiviewViewCount"`
	MeshOutputPerVertexGranularity        int   `json:"meshOutputPerVertexGranularity"`
	MeshOutputPerPrimitiveGranularity     int   `json:"meshOutputPerPrimitiveGranularity"`
	MaxPreferredTaskWorkGroupInvocations  int   `json:"maxPreferredTaskWorkGroupInvocations"`
	MaxPreferredMeshWorkGroupInvocations  int   `json:"maxPreferredMeshWorkGroupInvocations"`
	PrefersLocalInvocationVertexOutput    bool  `json:"prefersLocalInvocationVertexOutput"`
	PrefersLocalInvocationPrimitiveOutput bool  `json:"prefersLocalInvocationPrimitiveOutput"`
	PrefersCompactVertexOutput            bool  `json:"prefersCompactVertexOutput"`
	PrefersCompactPrimitiveOutput         bool  `json:"prefersCompactPrimitiveOutput"`
}

type VkPhysicalDeviceMultiDrawPropertiesEXT struct {
	MaxMultiDrawCount int `json:"maxMultiDrawCount"`
}

type VkPhysicalDeviceNestedCommandBufferPropertiesEXT struct {
	MaxCommandBufferNestingLevel int64 `json:"maxCommandBufferNestingLevel"`
}

type VkPhysicalDevicePCIBusInfoPropertiesEXT struct {
	PciDomain   int `json:"pciDomain"`
	PciBus      int `json:"pciBus"`
	PciDevice   int `json:"pciDevice"`
	PciFunction int `json:"pciFunction"`
}

type VkPhysicalDevicePipelineBinaryPropertiesKHR struct {
	PipelineBinaryInternalCache            bool `json:"pipelineBinaryInternalCache"`
	PipelineBinaryInternalCacheControl     bool `json:"pipelineBinaryInternalCacheControl"`
	PipelineBinaryPrefersInternalCache     bool `json:"pipelineBinaryPrefersInternalCache"`
	PipelineBinaryPrecompiledInternalCache bool `json:"pipelineBinaryPrecompiledInternalCache"`
	PipelineBinaryCompressedData           bool `json:"pipelineBinaryCompressedData"`
}

type VkPhysicalDeviceProvokingVertexPropertiesEXT struct {
	ProvokingVertexModePerPipeline                       bool `json:"provokingVertexModePerPipeline"`
	TransformFeedbackPreservesTriangleFanProvokingVertex bool `json:"transformFeedbackPreservesTriangleFanProvokingVertex"`
}

type VkPhysicalDeviceRayTracingPipelinePropertiesKHR struct {
	ShaderGroupHandleSize              int `json:"shaderGroupHandleSize"`
	MaxRayRecursionDepth               int `json:"maxRayRecursionDepth"`
	MaxShaderGroupStride               int `json:"maxShaderGroupStride"`
	ShaderGroupBaseAlignment           int `json:"shaderGroupBaseAlignment"`
	ShaderGroupHandleCaptureReplaySize int `json:"shaderGroupHandleCaptureReplaySize"`
	MaxRayDispatchInvocationCount      int `json:"maxRayDispatchInvocationCount"`
	ShaderGroupHandleAlignment         int `json:"shaderGroupHandleAlignment"`
	MaxRayHitAttributeSize             int `json:"maxRayHitAttributeSize"`
}

type VkPhysicalDeviceRobustness2PropertiesEXT struct {
	RobustStorageBufferAccessSizeAlignment int `json:"robustStorageBufferAccessSizeAlignment"`
	RobustUniformBufferAccessSizeAlignment int `json:"robustUniformBufferAccessSizeAlignment"`
}

type VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT struct {
	ShaderModuleIdentifierAlgorithmUUID []int `json:"shaderModuleIdentifierAlgorithmUUID"`
}

type VkPhysicalDeviceShaderObjectPropertiesEXT struct {
	ShaderBinaryUUID    []int `json:"shaderBinaryUUID"`
	ShaderBinaryVersion int   `json:"shaderBinaryVersion"`
}

type VkPhysicalDeviceTransformFeedbackPropertiesEXT struct {
	MaxTransformFeedbackStreams                int   `json:"maxTransformFeedbackStreams"`
	MaxTransformFeedbackBuffers                int   `json:"maxTransformFeedbackBuffers"`
	MaxTransformFeedbackBufferSize             int64 `json:"maxTransformFeedbackBufferSize"`
	MaxTransformFeedbackStreamDataSize         int   `json:"maxTransformFeedbackStreamDataSize"`
	MaxTransformFeedbackBufferDataSize         int   `json:"maxTransformFeedbackBufferDataSize"`
	MaxTransformFeedbackBufferDataStride       int   `json:"maxTransformFeedbackBufferDataStride"`
	TransformFeedbackQueries                   bool  `json:"transformFeedbackQueries"`
	TransformFeedbackStreamsLinesTriangles     bool  `json:"transformFeedbackStreamsLinesTriangles"`
	TransformFeedbackRasterizationStreamSelect bool  `json:"transformFeedbackRasterizationStreamSelect"`
	TransformFeedbackDraw                      bool  `json:"transformFeedbackDraw"`
}

type VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT struct {
	MaxVertexAttribDivisor int64 `json:"maxVertexAttribDivisor"`
}

type VkPhysicalDeviceVulkan11Properties struct {
	DeviceUUID                        []int    `json:"deviceUUID"`
	DriverUUID                        []int    `json:"driverUUID"`
	DeviceNodeMask                    int      `json:"deviceNodeMask"`
	DeviceLUIDValid                   bool     `json:"deviceLUIDValid"`
	SubgroupSize                      int      `json:"subgroupSize"`
	SubgroupSupportedStages           []string `json:"subgroupSupportedStages"`
	SubgroupSupportedOperations       []string `json:"subgroupSupportedOperations"`
	SubgroupQuadOperationsInAllStages bool     `json:"subgroupQuadOperationsInAllStages"`
	PointClippingBehavior             string   `json:"pointClippingBehavior"`
	MaxMultiviewViewCount             int      `json:"maxMultiviewViewCount"`
	MaxMultiviewInstanceIndex         int64    `json:"maxMultiviewInstanceIndex"`
	ProtectedNoFault                  bool     `json:"protectedNoFault"`
	MaxPerSetDescriptors              int      `json:"maxPerSetDescriptors"`
	MaxMemoryAllocationSize           int64    `json:"maxMemoryAllocationSize"`
}

type ConformanceVersion struct {
	Major    int `json:"major"`
	Minor    int `json:"minor"`
	Subminor int `json:"subminor"`
	Patch    int `json:"patch"`
}

type VkPhysicalDeviceVulkan12Properties struct {
	DriverID                                             string             `json:"driverID"`
	DriverName                                           string             `json:"driverName"`
	DriverInfo                                           string             `json:"driverInfo"`
	ConformanceVersion                                   ConformanceVersion `json:"conformanceVersion"`
	DenormBehaviorIndependence                           string             `json:"denormBehaviorIndependence"`
	RoundingModeIndependence                             string             `json:"roundingModeIndependence"`
	ShaderSignedZeroInfNanPreserveFloat16                bool               `json:"shaderSignedZeroInfNanPreserveFloat16"`
	ShaderSignedZeroInfNanPreserveFloat32                bool               `json:"shaderSignedZeroInfNanPreserveFloat32"`
	ShaderSignedZeroInfNanPreserveFloat64                bool               `json:"shaderSignedZeroInfNanPreserveFloat64"`
	ShaderDenormPreserveFloat16                          bool               `json:"shaderDenormPreserveFloat16"`
	ShaderDenormPreserveFloat32                          bool               `json:"shaderDenormPreserveFloat32"`
	ShaderDenormPreserveFloat64                          bool               `json:"shaderDenormPreserveFloat64"`
	ShaderDenormFlushToZeroFloat16                       bool               `json:"shaderDenormFlushToZeroFloat16"`
	ShaderDenormFlushToZeroFloat32                       bool               `json:"shaderDenormFlushToZeroFloat32"`
	ShaderDenormFlushToZeroFloat64                       bool               `json:"shaderDenormFlushToZeroFloat64"`
	ShaderRoundingModeRTEFloat16                         bool               `json:"shaderRoundingModeRTEFloat16"`
	ShaderRoundingModeRTEFloat32                         bool               `json:"shaderRoundingModeRTEFloat32"`
	ShaderRoundingModeRTEFloat64                         bool               `json:"shaderRoundingModeRTEFloat64"`
	ShaderRoundingModeRTZFloat16                         bool               `json:"shaderRoundingModeRTZFloat16"`
	ShaderRoundingModeRTZFloat32                         bool               `json:"shaderRoundingModeRTZFloat32"`
	ShaderRoundingModeRTZFloat64                         bool               `json:"shaderRoundingModeRTZFloat64"`
	MaxUpdateAfterBindDescriptorsInAllPools              int                `json:"maxUpdateAfterBindDescriptorsInAllPools"`
	ShaderUniformBufferArrayNonUniformIndexingNative     bool               `json:"shaderUniformBufferArrayNonUniformIndexingNative"`
	ShaderSampledImageArrayNonUniformIndexingNative      bool               `json:"shaderSampledImageArrayNonUniformIndexingNative"`
	ShaderStorageBufferArrayNonUniformIndexingNative     bool               `json:"shaderStorageBufferArrayNonUniformIndexingNative"`
	ShaderStorageImageArrayNonUniformIndexingNative      bool               `json:"shaderStorageImageArrayNonUniformIndexingNative"`
	ShaderInputAttachmentArrayNonUniformIndexingNative   bool               `json:"shaderInputAttachmentArrayNonUniformIndexingNative"`
	RobustBufferAccessUpdateAfterBind                    bool               `json:"robustBufferAccessUpdateAfterBind"`
	QuadDivergentImplicitLod                             bool               `json:"quadDivergentImplicitLod"`
	MaxPerStageDescriptorUpdateAfterBindSamplers         int                `json:"maxPerStageDescriptorUpdateAfterBindSamplers"`
	MaxPerStageDescriptorUpdateAfterBindUniformBuffers   int                `json:"maxPerStageDescriptorUpdateAfterBindUniformBuffers"`
	MaxPerStageDescriptorUpdateAfterBindStorageBuffers   int                `json:"maxPerStageDescriptorUpdateAfterBindStorageBuffers"`
	MaxPerStageDescriptorUpdateAfterBindSampledImages    int                `json:"maxPerStageDescriptorUpdateAfterBindSampledImages"`
	MaxPerStageDescriptorUpdateAfterBindStorageImages    int                `json:"maxPerStageDescriptorUpdateAfterBindStorageImages"`
	MaxPerStageDescriptorUpdateAfterBindInputAttachments int                `json:"maxPerStageDescriptorUpdateAfterBindInputAttachments"`
	MaxPerStageUpdateAfterBindResources                  int                `json:"maxPerStageUpdateAfterBindResources"`
	MaxDescriptorSetUpdateAfterBindSamplers              int                `json:"maxDescriptorSetUpdateAfterBindSamplers"`
	MaxDescriptorSetUpdateAfterBindUniformBuffers        int                `json:"maxDescriptorSetUpdateAfterBindUniformBuffers"`
	MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic int                `json:"maxDescriptorSetUpdateAfterBindUniformBuffersDynamic"`
	MaxDescriptorSetUpdateAfterBindStorageBuffers        int                `json:"maxDescriptorSetUpdateAfterBindStorageBuffers"`
	MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic int                `json:"maxDescriptorSetUpdateAfterBindStorageBuffersDynamic"`
	MaxDescriptorSetUpdateAfterBindSampledImages         int                `json:"maxDescriptorSetUpdateAfterBindSampledImages"`
	MaxDescriptorSetUpdateAfterBindStorageImages         int                `json:"maxDescriptorSetUpdateAfterBindStorageImages"`
	MaxDescriptorSetUpdateAfterBindInputAttachments      int                `json:"maxDescriptorSetUpdateAfterBindInputAttachments"`
	SupportedDepthResolveModes                           []string           `json:"supportedDepthResolveModes"`
	SupportedStencilResolveModes                         []string           `json:"supportedStencilResolveModes"`
	IndependentResolveNone                               bool               `json:"independentResolveNone"`
	IndependentResolve                                   bool               `json:"independentResolve"`
	FilterMinmaxSingleComponentFormats                   bool               `json:"filterMinmaxSingleComponentFormats"`
	FilterMinmaxImageComponentMapping                    bool               `json:"filterMinmaxImageComponentMapping"`
	MaxTimelineSemaphoreValueDifference                  int64              `json:"maxTimelineSemaphoreValueDifference"`
	FramebufferIntegerColorSampleCounts                  []string           `json:"framebufferIntegerColorSampleCounts"`
}

type VkPhysicalDeviceVulkan13Properties struct {
	MinSubgroupSize                                                               int      `json:"minSubgroupSize"`
	MaxSubgroupSize                                                               int      `json:"maxSubgroupSize"`
	MaxComputeWorkgroupSubgroups                                                  int64    `json:"maxComputeWorkgroupSubgroups"`
	RequiredSubgroupSizeStages                                                    []string `json:"requiredSubgroupSizeStages"`
	MaxInlineUniformBlockSize                                                     int      `json:"maxInlineUniformBlockSize"`
	MaxPerStageDescriptorInlineUniformBlocks                                      int      `json:"maxPerStageDescriptorInlineUniformBlocks"`
	MaxPerStageDescriptorUpdateAfterBindInlineUniformBlocks                       int      `json:"maxPerStageDescriptorUpdateAfterBindInlineUniformBlocks"`
	MaxDescriptorSetInlineUniformBlocks                                           int      `json:"maxDescriptorSetInlineUniformBlocks"`
	MaxDescriptorSetUpdateAfterBindInlineUniformBlocks                            int      `json:"maxDescriptorSetUpdateAfterBindInlineUniformBlocks"`
	MaxInlineUniformTotalSize                                                     int      `json:"maxInlineUniformTotalSize"`
	IntegerDotProduct8BitUnsignedAccelerated                                      bool     `json:"integerDotProduct8BitUnsignedAccelerated"`
	IntegerDotProduct8BitSignedAccelerated                                        bool     `json:"integerDotProduct8BitSignedAccelerated"`
	IntegerDotProduct8BitMixedSignednessAccelerated                               bool     `json:"integerDotProduct8BitMixedSignednessAccelerated"`
	IntegerDotProduct4X8BitPackedUnsignedAccelerated                              bool     `json:"integerDotProduct4x8BitPackedUnsignedAccelerated"`
	IntegerDotProduct4X8BitPackedSignedAccelerated                                bool     `json:"integerDotProduct4x8BitPackedSignedAccelerated"`
	IntegerDotProduct4X8BitPackedMixedSignednessAccelerated                       bool     `json:"integerDotProduct4x8BitPackedMixedSignednessAccelerated"`
	IntegerDotProduct16BitUnsignedAccelerated                                     bool     `json:"integerDotProduct16BitUnsignedAccelerated"`
	IntegerDotProduct16BitSignedAccelerated                                       bool     `json:"integerDotProduct16BitSignedAccelerated"`
	IntegerDotProduct16BitMixedSignednessAccelerated                              bool     `json:"integerDotProduct16BitMixedSignednessAccelerated"`
	IntegerDotProduct32BitUnsignedAccelerated                                     bool     `json:"integerDotProduct32BitUnsignedAccelerated"`
	IntegerDotProduct32BitSignedAccelerated                                       bool     `json:"integerDotProduct32BitSignedAccelerated"`
	IntegerDotProduct32BitMixedSignednessAccelerated                              bool     `json:"integerDotProduct32BitMixedSignednessAccelerated"`
	IntegerDotProduct64BitUnsignedAccelerated                                     bool     `json:"integerDotProduct64BitUnsignedAccelerated"`
	IntegerDotProduct64BitSignedAccelerated                                       bool     `json:"integerDotProduct64BitSignedAccelerated"`
	IntegerDotProduct64BitMixedSignednessAccelerated                              bool     `json:"integerDotProduct64BitMixedSignednessAccelerated"`
	IntegerDotProductAccumulatingSaturating8BitUnsignedAccelerated                bool     `json:"integerDotProductAccumulatingSaturating8BitUnsignedAccelerated"`
	IntegerDotProductAccumulatingSaturating8BitSignedAccelerated                  bool     `json:"integerDotProductAccumulatingSaturating8BitSignedAccelerated"`
	IntegerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated         bool     `json:"integerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated"`
	IntegerDotProductAccumulatingSaturating4X8BitPackedUnsignedAccelerated        bool     `json:"integerDotProductAccumulatingSaturating4x8BitPackedUnsignedAccelerated"`
	IntegerDotProductAccumulatingSaturating4X8BitPackedSignedAccelerated          bool     `json:"integerDotProductAccumulatingSaturating4x8BitPackedSignedAccelerated"`
	IntegerDotProductAccumulatingSaturating4X8BitPackedMixedSignednessAccelerated bool     `json:"integerDotProductAccumulatingSaturating4x8BitPackedMixedSignednessAccelerated"`
	IntegerDotProductAccumulatingSaturating16BitUnsignedAccelerated               bool     `json:"integerDotProductAccumulatingSaturating16BitUnsignedAccelerated"`
	IntegerDotProductAccumulatingSaturating16BitSignedAccelerated                 bool     `json:"integerDotProductAccumulatingSaturating16BitSignedAccelerated"`
	IntegerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated        bool     `json:"integerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated"`
	IntegerDotProductAccumulatingSaturating32BitUnsignedAccelerated               bool     `json:"integerDotProductAccumulatingSaturating32BitUnsignedAccelerated"`
	IntegerDotProductAccumulatingSaturating32BitSignedAccelerated                 bool     `json:"integerDotProductAccumulatingSaturating32BitSignedAccelerated"`
	IntegerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated        bool     `json:"integerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated"`
	IntegerDotProductAccumulatingSaturating64BitUnsignedAccelerated               bool     `json:"integerDotProductAccumulatingSaturating64BitUnsignedAccelerated"`
	IntegerDotProductAccumulatingSaturating64BitSignedAccelerated                 bool     `json:"integerDotProductAccumulatingSaturating64BitSignedAccelerated"`
	IntegerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated        bool     `json:"integerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated"`
	StorageTexelBufferOffsetAlignmentBytes                                        int      `json:"storageTexelBufferOffsetAlignmentBytes"`
	StorageTexelBufferOffsetSingleTexelAlignment                                  bool     `json:"storageTexelBufferOffsetSingleTexelAlignment"`
	UniformTexelBufferOffsetAlignmentBytes                                        int      `json:"uniformTexelBufferOffsetAlignmentBytes"`
	UniformTexelBufferOffsetSingleTexelAlignment                                  bool     `json:"uniformTexelBufferOffsetSingleTexelAlignment"`
	MaxBufferSize                                                                 int64    `json:"maxBufferSize"`
}

type VkPhysicalDeviceVulkan14Properties struct {
	LineSubPixelPrecisionBits                           int    `json:"lineSubPixelPrecisionBits"`
	MaxVertexAttribDivisor                              int64  `json:"maxVertexAttribDivisor"`
	SupportsNonZeroFirstInstance                        bool   `json:"supportsNonZeroFirstInstance"`
	MaxPushDescriptors                                  int    `json:"maxPushDescriptors"`
	DynamicRenderingLocalReadDepthStencilAttachments    bool   `json:"dynamicRenderingLocalReadDepthStencilAttachments"`
	DynamicRenderingLocalReadMultisampledAttachments    bool   `json:"dynamicRenderingLocalReadMultisampledAttachments"`
	EarlyFragmentMultisampleCoverageAfterSampleCounting bool   `json:"earlyFragmentMultisampleCoverageAfterSampleCounting"`
	EarlyFragmentSampleMaskTestBeforeSampleCounting     bool   `json:"earlyFragmentSampleMaskTestBeforeSampleCounting"`
	DepthStencilSwizzleOneSupport                       bool   `json:"depthStencilSwizzleOneSupport"`
	PolygonModePointSize                                bool   `json:"polygonModePointSize"`
	NonStrictSinglePixelWideLinesUseParallelogram       bool   `json:"nonStrictSinglePixelWideLinesUseParallelogram"`
	NonStrictWideLinesUseParallelogram                  bool   `json:"nonStrictWideLinesUseParallelogram"`
	BlockTexelViewCompatibleMultipleLayers              bool   `json:"blockTexelViewCompatibleMultipleLayers"`
	MaxCombinedImageSamplerDescriptorCount              int    `json:"maxCombinedImageSamplerDescriptorCount"`
	FragmentShadingRateClampCombinerInputs              bool   `json:"fragmentShadingRateClampCombinerInputs"`
	DefaultRobustnessStorageBuffers                     string `json:"defaultRobustnessStorageBuffers"`
	DefaultRobustnessUniformBuffers                     string `json:"defaultRobustnessUniformBuffers"`
	DefaultRobustnessVertexInputs                       string `json:"defaultRobustnessVertexInputs"`
	DefaultRobustnessImages                             string `json:"defaultRobustnessImages"`
	CopySrcLayoutCount                                  int    `json:"copySrcLayoutCount"`
	PCopySrcLayouts                                     string `json:"pCopySrcLayouts"`
	CopyDstLayoutCount                                  int    `json:"copyDstLayoutCount"`
	PCopyDstLayouts                                     string `json:"pCopyDstLayouts"`
	OptimalTilingLayoutUUID                             []int  `json:"optimalTilingLayoutUUID"`
	IdenticalMemoryTypeRequirements                     bool   `json:"identicalMemoryTypeRequirements"`
}

type Properties struct {
	VkPhysicalDeviceProperties                              VkPhysicalDeviceProperties                              `json:"VkPhysicalDeviceProperties"`
	VkPhysicalDeviceAccelerationStructurePropertiesKHR      VkPhysicalDeviceAccelerationStructurePropertiesKHR      `json:"VkPhysicalDeviceAccelerationStructurePropertiesKHR"`
	VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR   VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR   `json:"VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR"`
	VkPhysicalDeviceConservativeRasterizationPropertiesEXT  VkPhysicalDeviceConservativeRasterizationPropertiesEXT  `json:"VkPhysicalDeviceConservativeRasterizationPropertiesEXT"`
	VkPhysicalDeviceCooperativeMatrixPropertiesKHR          VkPhysicalDeviceCooperativeMatrixPropertiesKHR          `json:"VkPhysicalDeviceCooperativeMatrixPropertiesKHR"`
	VkPhysicalDeviceCustomBorderColorPropertiesEXT          VkPhysicalDeviceCustomBorderColorPropertiesEXT          `json:"VkPhysicalDeviceCustomBorderColorPropertiesEXT"`
	VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT `json:"VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT"`
	VkPhysicalDeviceDescriptorBufferPropertiesEXT           VkPhysicalDeviceDescriptorBufferPropertiesEXT           `json:"VkPhysicalDeviceDescriptorBufferPropertiesEXT"`
	VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT    VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT    `json:"VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT"`
	VkPhysicalDeviceDiscardRectanglePropertiesEXT           VkPhysicalDeviceDiscardRectanglePropertiesEXT           `json:"VkPhysicalDeviceDiscardRectanglePropertiesEXT"`
	VkPhysicalDeviceDrmPropertiesEXT                        VkPhysicalDeviceDrmPropertiesEXT                        `json:"VkPhysicalDeviceDrmPropertiesEXT"`
	VkPhysicalDeviceExtendedDynamicState3PropertiesEXT      VkPhysicalDeviceExtendedDynamicState3PropertiesEXT      `json:"VkPhysicalDeviceExtendedDynamicState3PropertiesEXT"`
	VkPhysicalDeviceExternalMemoryHostPropertiesEXT         VkPhysicalDeviceExternalMemoryHostPropertiesEXT         `json:"VkPhysicalDeviceExternalMemoryHostPropertiesEXT"`
	VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR  VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR  `json:"VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR"`
	VkPhysicalDeviceFragmentShadingRatePropertiesKHR        VkPhysicalDeviceFragmentShadingRatePropertiesKHR        `json:"VkPhysicalDeviceFragmentShadingRatePropertiesKHR"`
	VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT    VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT    `json:"VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT"`
	VkPhysicalDeviceLayeredAPIPropertiesListKHR             VkPhysicalDeviceLayeredAPIPropertiesListKHR             `json:"VkPhysicalDeviceLayeredApiPropertiesListKHR"`
	VkPhysicalDeviceLegacyVertexAttributesPropertiesEXT     VkPhysicalDeviceLegacyVertexAttributesPropertiesEXT     `json:"VkPhysicalDeviceLegacyVertexAttributesPropertiesEXT"`
	VkPhysicalDeviceMaintenance7PropertiesKHR               VkPhysicalDeviceMaintenance7PropertiesKHR               `json:"VkPhysicalDeviceMaintenance7PropertiesKHR"`
	VkPhysicalDeviceMapMemoryPlacedPropertiesEXT            VkPhysicalDeviceMapMemoryPlacedPropertiesEXT            `json:"VkPhysicalDeviceMapMemoryPlacedPropertiesEXT"`
	VkPhysicalDeviceMeshShaderPropertiesEXT                 VkPhysicalDeviceMeshShaderPropertiesEXT                 `json:"VkPhysicalDeviceMeshShaderPropertiesEXT"`
	VkPhysicalDeviceMultiDrawPropertiesEXT                  VkPhysicalDeviceMultiDrawPropertiesEXT                  `json:"VkPhysicalDeviceMultiDrawPropertiesEXT"`
	VkPhysicalDeviceNestedCommandBufferPropertiesEXT        VkPhysicalDeviceNestedCommandBufferPropertiesEXT        `json:"VkPhysicalDeviceNestedCommandBufferPropertiesEXT"`
	VkPhysicalDevicePCIBusInfoPropertiesEXT                 VkPhysicalDevicePCIBusInfoPropertiesEXT                 `json:"VkPhysicalDevicePCIBusInfoPropertiesEXT"`
	VkPhysicalDevicePipelineBinaryPropertiesKHR             VkPhysicalDevicePipelineBinaryPropertiesKHR             `json:"VkPhysicalDevicePipelineBinaryPropertiesKHR"`
	VkPhysicalDeviceProvokingVertexPropertiesEXT            VkPhysicalDeviceProvokingVertexPropertiesEXT            `json:"VkPhysicalDeviceProvokingVertexPropertiesEXT"`
	VkPhysicalDeviceRayTracingPipelinePropertiesKHR         VkPhysicalDeviceRayTracingPipelinePropertiesKHR         `json:"VkPhysicalDeviceRayTracingPipelinePropertiesKHR"`
	VkPhysicalDeviceRobustness2PropertiesEXT                VkPhysicalDeviceRobustness2PropertiesEXT                `json:"VkPhysicalDeviceRobustness2PropertiesEXT"`
	VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT     VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT     `json:"VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT"`
	VkPhysicalDeviceShaderObjectPropertiesEXT               VkPhysicalDeviceShaderObjectPropertiesEXT               `json:"VkPhysicalDeviceShaderObjectPropertiesEXT"`
	VkPhysicalDeviceTransformFeedbackPropertiesEXT          VkPhysicalDeviceTransformFeedbackPropertiesEXT          `json:"VkPhysicalDeviceTransformFeedbackPropertiesEXT"`
	VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT     VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT     `json:"VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT"`
	VkPhysicalDeviceVulkan11Properties                      VkPhysicalDeviceVulkan11Properties                      `json:"VkPhysicalDeviceVulkan11Properties"`
	VkPhysicalDeviceVulkan12Properties                      VkPhysicalDeviceVulkan12Properties                      `json:"VkPhysicalDeviceVulkan12Properties"`
	VkPhysicalDeviceVulkan13Properties                      VkPhysicalDeviceVulkan13Properties                      `json:"VkPhysicalDeviceVulkan13Properties"`
	VkPhysicalDeviceVulkan14Properties                      VkPhysicalDeviceVulkan14Properties                      `json:"VkPhysicalDeviceVulkan14Properties"`
}

type VkFormatProperties struct {
	LinearTilingFeatures  []string `json:"linearTilingFeatures"`
	OptimalTilingFeatures []string `json:"optimalTilingFeatures"`
	BufferFeatures        []any    `json:"bufferFeatures"`
}

type VkFormatR4G4B4A4UnormPack16 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB4G4R4A4UnormPack16 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR5G6B5UnormPack16 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB5G6R5UnormPack16 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR5G5B5A1UnormPack16 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB5G5R5A1UnormPack16 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA1R5G5B5UnormPack16 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8Snorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8Uscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8Sscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8Srgb struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8Snorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8Uscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8Sscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8Srgb struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8B8Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8B8Snorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8B8Uscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8B8Sscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8B8Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8B8Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB8G8R8Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB8G8R8Snorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB8G8R8Uscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB8G8R8Sscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB8G8R8Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB8G8R8Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8B8A8Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8B8A8Snorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8B8A8Uscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8B8A8Sscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8B8A8Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8B8A8Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR8G8B8A8Srgb struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB8G8R8A8Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB8G8R8A8Snorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB8G8R8A8Uscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB8G8R8A8Sscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB8G8R8A8Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB8G8R8A8Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB8G8R8A8Srgb struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA8B8G8R8UnormPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA8B8G8R8SnormPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA8B8G8R8UscaledPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA8B8G8R8SscaledPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA8B8G8R8UintPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA8B8G8R8SintPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA8B8G8R8SrgbPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA2R10G10B10UnormPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA2R10G10B10SnormPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA2R10G10B10UscaledPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA2R10G10B10SscaledPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA2R10G10B10UintPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA2R10G10B10SintPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA2B10G10R10UnormPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA2B10G10R10SnormPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA2B10G10R10UscaledPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA2B10G10R10SscaledPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA2B10G10R10UintPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA2B10G10R10SintPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16Snorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16Uscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16Sscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16Sfloat struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16Snorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16Uscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16Sscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16Sfloat struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16B16Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16B16Snorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16B16Uscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16B16Sscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16B16Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16B16Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16B16Sfloat struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16B16A16Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16B16A16Snorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16B16A16Uscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16B16A16Sscaled struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16B16A16Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16B16A16Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR16G16B16A16Sfloat struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR32Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR32Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR32Sfloat struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR32G32Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR32G32Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR32G32Sfloat struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR32G32B32Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR32G32B32Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR32G32B32Sfloat struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR32G32B32A32Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR32G32B32A32Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR32G32B32A32Sfloat struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR64Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR64Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR64Sfloat struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR64G64Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR64G64Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR64G64Sfloat struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR64G64B64Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR64G64B64Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR64G64B64Sfloat struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR64G64B64A64Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR64G64B64A64Sint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR64G64B64A64Sfloat struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatB10G11R11UfloatPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatE5B9G9R9UfloatPack32 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatD16Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatD32Sfloat struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatS8Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatD16UnormS8Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatD32SfloatS8Uint struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatBc1RgbUnormBlock struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatBc1RgbSrgbBlock struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatBc1RgbaUnormBlock struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatBc1RgbaSrgbBlock struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatBc2UnormBlock struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatBc2SrgbBlock struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatBc3UnormBlock struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatBc3SrgbBlock struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatBc4UnormBlock struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatBc4SnormBlock struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatBc5UnormBlock struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatBc5SnormBlock struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatBc6HUfloatBlock struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatBc6HSfloatBlock struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatBc7UnormBlock struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatBc7SrgbBlock struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatG8B8R83Plane420Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatG8B8R82Plane420Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatG8B8R83Plane422Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatG8B8R82Plane422Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatG8B8R83Plane444Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR10X6UnormPack16 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR10X6G10X6Unorm2Pack16 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatG10X6B10X6R10X62Plane420Unorm3Pack16 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR12X4UnormPack16 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatR12X4G12X4Unorm2Pack16 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatG12X4B12X4R12X42Plane420Unorm3Pack16 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatG16B16R163Plane420Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatG16B16R162Plane420Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatG16B16R163Plane422Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatG16B16R162Plane422Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatG16B16R163Plane444Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA4R4G4B4UnormPack16 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA4B4G4R4UnormPack16 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA1B5G5R5UnormPack16 struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type VkFormatA8Unorm struct {
	VkFormatProperties VkFormatProperties `json:"VkFormatProperties"`
}

type Formats struct {
	VkFormatR4G4B4A4UnormPack16                  VkFormatR4G4B4A4UnormPack16                  `json:"VK_FORMAT_R4G4B4A4_UNORM_PACK16"`
	VkFormatB4G4R4A4UnormPack16                  VkFormatB4G4R4A4UnormPack16                  `json:"VK_FORMAT_B4G4R4A4_UNORM_PACK16"`
	VkFormatR5G6B5UnormPack16                    VkFormatR5G6B5UnormPack16                    `json:"VK_FORMAT_R5G6B5_UNORM_PACK16"`
	VkFormatB5G6R5UnormPack16                    VkFormatB5G6R5UnormPack16                    `json:"VK_FORMAT_B5G6R5_UNORM_PACK16"`
	VkFormatR5G5B5A1UnormPack16                  VkFormatR5G5B5A1UnormPack16                  `json:"VK_FORMAT_R5G5B5A1_UNORM_PACK16"`
	VkFormatB5G5R5A1UnormPack16                  VkFormatB5G5R5A1UnormPack16                  `json:"VK_FORMAT_B5G5R5A1_UNORM_PACK16"`
	VkFormatA1R5G5B5UnormPack16                  VkFormatA1R5G5B5UnormPack16                  `json:"VK_FORMAT_A1R5G5B5_UNORM_PACK16"`
	VkFormatR8Unorm                              VkFormatR8Unorm                              `json:"VK_FORMAT_R8_UNORM"`
	VkFormatR8Snorm                              VkFormatR8Snorm                              `json:"VK_FORMAT_R8_SNORM"`
	VkFormatR8Uscaled                            VkFormatR8Uscaled                            `json:"VK_FORMAT_R8_USCALED"`
	VkFormatR8Sscaled                            VkFormatR8Sscaled                            `json:"VK_FORMAT_R8_SSCALED"`
	VkFormatR8Uint                               VkFormatR8Uint                               `json:"VK_FORMAT_R8_UINT"`
	VkFormatR8Sint                               VkFormatR8Sint                               `json:"VK_FORMAT_R8_SINT"`
	VkFormatR8Srgb                               VkFormatR8Srgb                               `json:"VK_FORMAT_R8_SRGB"`
	VkFormatR8G8Unorm                            VkFormatR8G8Unorm                            `json:"VK_FORMAT_R8G8_UNORM"`
	VkFormatR8G8Snorm                            VkFormatR8G8Snorm                            `json:"VK_FORMAT_R8G8_SNORM"`
	VkFormatR8G8Uscaled                          VkFormatR8G8Uscaled                          `json:"VK_FORMAT_R8G8_USCALED"`
	VkFormatR8G8Sscaled                          VkFormatR8G8Sscaled                          `json:"VK_FORMAT_R8G8_SSCALED"`
	VkFormatR8G8Uint                             VkFormatR8G8Uint                             `json:"VK_FORMAT_R8G8_UINT"`
	VkFormatR8G8Sint                             VkFormatR8G8Sint                             `json:"VK_FORMAT_R8G8_SINT"`
	VkFormatR8G8Srgb                             VkFormatR8G8Srgb                             `json:"VK_FORMAT_R8G8_SRGB"`
	VkFormatR8G8B8Unorm                          VkFormatR8G8B8Unorm                          `json:"VK_FORMAT_R8G8B8_UNORM"`
	VkFormatR8G8B8Snorm                          VkFormatR8G8B8Snorm                          `json:"VK_FORMAT_R8G8B8_SNORM"`
	VkFormatR8G8B8Uscaled                        VkFormatR8G8B8Uscaled                        `json:"VK_FORMAT_R8G8B8_USCALED"`
	VkFormatR8G8B8Sscaled                        VkFormatR8G8B8Sscaled                        `json:"VK_FORMAT_R8G8B8_SSCALED"`
	VkFormatR8G8B8Uint                           VkFormatR8G8B8Uint                           `json:"VK_FORMAT_R8G8B8_UINT"`
	VkFormatR8G8B8Sint                           VkFormatR8G8B8Sint                           `json:"VK_FORMAT_R8G8B8_SINT"`
	VkFormatB8G8R8Unorm                          VkFormatB8G8R8Unorm                          `json:"VK_FORMAT_B8G8R8_UNORM"`
	VkFormatB8G8R8Snorm                          VkFormatB8G8R8Snorm                          `json:"VK_FORMAT_B8G8R8_SNORM"`
	VkFormatB8G8R8Uscaled                        VkFormatB8G8R8Uscaled                        `json:"VK_FORMAT_B8G8R8_USCALED"`
	VkFormatB8G8R8Sscaled                        VkFormatB8G8R8Sscaled                        `json:"VK_FORMAT_B8G8R8_SSCALED"`
	VkFormatB8G8R8Uint                           VkFormatB8G8R8Uint                           `json:"VK_FORMAT_B8G8R8_UINT"`
	VkFormatB8G8R8Sint                           VkFormatB8G8R8Sint                           `json:"VK_FORMAT_B8G8R8_SINT"`
	VkFormatR8G8B8A8Unorm                        VkFormatR8G8B8A8Unorm                        `json:"VK_FORMAT_R8G8B8A8_UNORM"`
	VkFormatR8G8B8A8Snorm                        VkFormatR8G8B8A8Snorm                        `json:"VK_FORMAT_R8G8B8A8_SNORM"`
	VkFormatR8G8B8A8Uscaled                      VkFormatR8G8B8A8Uscaled                      `json:"VK_FORMAT_R8G8B8A8_USCALED"`
	VkFormatR8G8B8A8Sscaled                      VkFormatR8G8B8A8Sscaled                      `json:"VK_FORMAT_R8G8B8A8_SSCALED"`
	VkFormatR8G8B8A8Uint                         VkFormatR8G8B8A8Uint                         `json:"VK_FORMAT_R8G8B8A8_UINT"`
	VkFormatR8G8B8A8Sint                         VkFormatR8G8B8A8Sint                         `json:"VK_FORMAT_R8G8B8A8_SINT"`
	VkFormatR8G8B8A8Srgb                         VkFormatR8G8B8A8Srgb                         `json:"VK_FORMAT_R8G8B8A8_SRGB"`
	VkFormatB8G8R8A8Unorm                        VkFormatB8G8R8A8Unorm                        `json:"VK_FORMAT_B8G8R8A8_UNORM"`
	VkFormatB8G8R8A8Snorm                        VkFormatB8G8R8A8Snorm                        `json:"VK_FORMAT_B8G8R8A8_SNORM"`
	VkFormatB8G8R8A8Uscaled                      VkFormatB8G8R8A8Uscaled                      `json:"VK_FORMAT_B8G8R8A8_USCALED"`
	VkFormatB8G8R8A8Sscaled                      VkFormatB8G8R8A8Sscaled                      `json:"VK_FORMAT_B8G8R8A8_SSCALED"`
	VkFormatB8G8R8A8Uint                         VkFormatB8G8R8A8Uint                         `json:"VK_FORMAT_B8G8R8A8_UINT"`
	VkFormatB8G8R8A8Sint                         VkFormatB8G8R8A8Sint                         `json:"VK_FORMAT_B8G8R8A8_SINT"`
	VkFormatB8G8R8A8Srgb                         VkFormatB8G8R8A8Srgb                         `json:"VK_FORMAT_B8G8R8A8_SRGB"`
	VkFormatA8B8G8R8UnormPack32                  VkFormatA8B8G8R8UnormPack32                  `json:"VK_FORMAT_A8B8G8R8_UNORM_PACK32"`
	VkFormatA8B8G8R8SnormPack32                  VkFormatA8B8G8R8SnormPack32                  `json:"VK_FORMAT_A8B8G8R8_SNORM_PACK32"`
	VkFormatA8B8G8R8UscaledPack32                VkFormatA8B8G8R8UscaledPack32                `json:"VK_FORMAT_A8B8G8R8_USCALED_PACK32"`
	VkFormatA8B8G8R8SscaledPack32                VkFormatA8B8G8R8SscaledPack32                `json:"VK_FORMAT_A8B8G8R8_SSCALED_PACK32"`
	VkFormatA8B8G8R8UintPack32                   VkFormatA8B8G8R8UintPack32                   `json:"VK_FORMAT_A8B8G8R8_UINT_PACK32"`
	VkFormatA8B8G8R8SintPack32                   VkFormatA8B8G8R8SintPack32                   `json:"VK_FORMAT_A8B8G8R8_SINT_PACK32"`
	VkFormatA8B8G8R8SrgbPack32                   VkFormatA8B8G8R8SrgbPack32                   `json:"VK_FORMAT_A8B8G8R8_SRGB_PACK32"`
	VkFormatA2R10G10B10UnormPack32               VkFormatA2R10G10B10UnormPack32               `json:"VK_FORMAT_A2R10G10B10_UNORM_PACK32"`
	VkFormatA2R10G10B10SnormPack32               VkFormatA2R10G10B10SnormPack32               `json:"VK_FORMAT_A2R10G10B10_SNORM_PACK32"`
	VkFormatA2R10G10B10UscaledPack32             VkFormatA2R10G10B10UscaledPack32             `json:"VK_FORMAT_A2R10G10B10_USCALED_PACK32"`
	VkFormatA2R10G10B10SscaledPack32             VkFormatA2R10G10B10SscaledPack32             `json:"VK_FORMAT_A2R10G10B10_SSCALED_PACK32"`
	VkFormatA2R10G10B10UintPack32                VkFormatA2R10G10B10UintPack32                `json:"VK_FORMAT_A2R10G10B10_UINT_PACK32"`
	VkFormatA2R10G10B10SintPack32                VkFormatA2R10G10B10SintPack32                `json:"VK_FORMAT_A2R10G10B10_SINT_PACK32"`
	VkFormatA2B10G10R10UnormPack32               VkFormatA2B10G10R10UnormPack32               `json:"VK_FORMAT_A2B10G10R10_UNORM_PACK32"`
	VkFormatA2B10G10R10SnormPack32               VkFormatA2B10G10R10SnormPack32               `json:"VK_FORMAT_A2B10G10R10_SNORM_PACK32"`
	VkFormatA2B10G10R10UscaledPack32             VkFormatA2B10G10R10UscaledPack32             `json:"VK_FORMAT_A2B10G10R10_USCALED_PACK32"`
	VkFormatA2B10G10R10SscaledPack32             VkFormatA2B10G10R10SscaledPack32             `json:"VK_FORMAT_A2B10G10R10_SSCALED_PACK32"`
	VkFormatA2B10G10R10UintPack32                VkFormatA2B10G10R10UintPack32                `json:"VK_FORMAT_A2B10G10R10_UINT_PACK32"`
	VkFormatA2B10G10R10SintPack32                VkFormatA2B10G10R10SintPack32                `json:"VK_FORMAT_A2B10G10R10_SINT_PACK32"`
	VkFormatR16Unorm                             VkFormatR16Unorm                             `json:"VK_FORMAT_R16_UNORM"`
	VkFormatR16Snorm                             VkFormatR16Snorm                             `json:"VK_FORMAT_R16_SNORM"`
	VkFormatR16Uscaled                           VkFormatR16Uscaled                           `json:"VK_FORMAT_R16_USCALED"`
	VkFormatR16Sscaled                           VkFormatR16Sscaled                           `json:"VK_FORMAT_R16_SSCALED"`
	VkFormatR16Uint                              VkFormatR16Uint                              `json:"VK_FORMAT_R16_UINT"`
	VkFormatR16Sint                              VkFormatR16Sint                              `json:"VK_FORMAT_R16_SINT"`
	VkFormatR16Sfloat                            VkFormatR16Sfloat                            `json:"VK_FORMAT_R16_SFLOAT"`
	VkFormatR16G16Unorm                          VkFormatR16G16Unorm                          `json:"VK_FORMAT_R16G16_UNORM"`
	VkFormatR16G16Snorm                          VkFormatR16G16Snorm                          `json:"VK_FORMAT_R16G16_SNORM"`
	VkFormatR16G16Uscaled                        VkFormatR16G16Uscaled                        `json:"VK_FORMAT_R16G16_USCALED"`
	VkFormatR16G16Sscaled                        VkFormatR16G16Sscaled                        `json:"VK_FORMAT_R16G16_SSCALED"`
	VkFormatR16G16Uint                           VkFormatR16G16Uint                           `json:"VK_FORMAT_R16G16_UINT"`
	VkFormatR16G16Sint                           VkFormatR16G16Sint                           `json:"VK_FORMAT_R16G16_SINT"`
	VkFormatR16G16Sfloat                         VkFormatR16G16Sfloat                         `json:"VK_FORMAT_R16G16_SFLOAT"`
	VkFormatR16G16B16Unorm                       VkFormatR16G16B16Unorm                       `json:"VK_FORMAT_R16G16B16_UNORM"`
	VkFormatR16G16B16Snorm                       VkFormatR16G16B16Snorm                       `json:"VK_FORMAT_R16G16B16_SNORM"`
	VkFormatR16G16B16Uscaled                     VkFormatR16G16B16Uscaled                     `json:"VK_FORMAT_R16G16B16_USCALED"`
	VkFormatR16G16B16Sscaled                     VkFormatR16G16B16Sscaled                     `json:"VK_FORMAT_R16G16B16_SSCALED"`
	VkFormatR16G16B16Uint                        VkFormatR16G16B16Uint                        `json:"VK_FORMAT_R16G16B16_UINT"`
	VkFormatR16G16B16Sint                        VkFormatR16G16B16Sint                        `json:"VK_FORMAT_R16G16B16_SINT"`
	VkFormatR16G16B16Sfloat                      VkFormatR16G16B16Sfloat                      `json:"VK_FORMAT_R16G16B16_SFLOAT"`
	VkFormatR16G16B16A16Unorm                    VkFormatR16G16B16A16Unorm                    `json:"VK_FORMAT_R16G16B16A16_UNORM"`
	VkFormatR16G16B16A16Snorm                    VkFormatR16G16B16A16Snorm                    `json:"VK_FORMAT_R16G16B16A16_SNORM"`
	VkFormatR16G16B16A16Uscaled                  VkFormatR16G16B16A16Uscaled                  `json:"VK_FORMAT_R16G16B16A16_USCALED"`
	VkFormatR16G16B16A16Sscaled                  VkFormatR16G16B16A16Sscaled                  `json:"VK_FORMAT_R16G16B16A16_SSCALED"`
	VkFormatR16G16B16A16Uint                     VkFormatR16G16B16A16Uint                     `json:"VK_FORMAT_R16G16B16A16_UINT"`
	VkFormatR16G16B16A16Sint                     VkFormatR16G16B16A16Sint                     `json:"VK_FORMAT_R16G16B16A16_SINT"`
	VkFormatR16G16B16A16Sfloat                   VkFormatR16G16B16A16Sfloat                   `json:"VK_FORMAT_R16G16B16A16_SFLOAT"`
	VkFormatR32Uint                              VkFormatR32Uint                              `json:"VK_FORMAT_R32_UINT"`
	VkFormatR32Sint                              VkFormatR32Sint                              `json:"VK_FORMAT_R32_SINT"`
	VkFormatR32Sfloat                            VkFormatR32Sfloat                            `json:"VK_FORMAT_R32_SFLOAT"`
	VkFormatR32G32Uint                           VkFormatR32G32Uint                           `json:"VK_FORMAT_R32G32_UINT"`
	VkFormatR32G32Sint                           VkFormatR32G32Sint                           `json:"VK_FORMAT_R32G32_SINT"`
	VkFormatR32G32Sfloat                         VkFormatR32G32Sfloat                         `json:"VK_FORMAT_R32G32_SFLOAT"`
	VkFormatR32G32B32Uint                        VkFormatR32G32B32Uint                        `json:"VK_FORMAT_R32G32B32_UINT"`
	VkFormatR32G32B32Sint                        VkFormatR32G32B32Sint                        `json:"VK_FORMAT_R32G32B32_SINT"`
	VkFormatR32G32B32Sfloat                      VkFormatR32G32B32Sfloat                      `json:"VK_FORMAT_R32G32B32_SFLOAT"`
	VkFormatR32G32B32A32Uint                     VkFormatR32G32B32A32Uint                     `json:"VK_FORMAT_R32G32B32A32_UINT"`
	VkFormatR32G32B32A32Sint                     VkFormatR32G32B32A32Sint                     `json:"VK_FORMAT_R32G32B32A32_SINT"`
	VkFormatR32G32B32A32Sfloat                   VkFormatR32G32B32A32Sfloat                   `json:"VK_FORMAT_R32G32B32A32_SFLOAT"`
	VkFormatR64Uint                              VkFormatR64Uint                              `json:"VK_FORMAT_R64_UINT"`
	VkFormatR64Sint                              VkFormatR64Sint                              `json:"VK_FORMAT_R64_SINT"`
	VkFormatR64Sfloat                            VkFormatR64Sfloat                            `json:"VK_FORMAT_R64_SFLOAT"`
	VkFormatR64G64Uint                           VkFormatR64G64Uint                           `json:"VK_FORMAT_R64G64_UINT"`
	VkFormatR64G64Sint                           VkFormatR64G64Sint                           `json:"VK_FORMAT_R64G64_SINT"`
	VkFormatR64G64Sfloat                         VkFormatR64G64Sfloat                         `json:"VK_FORMAT_R64G64_SFLOAT"`
	VkFormatR64G64B64Uint                        VkFormatR64G64B64Uint                        `json:"VK_FORMAT_R64G64B64_UINT"`
	VkFormatR64G64B64Sint                        VkFormatR64G64B64Sint                        `json:"VK_FORMAT_R64G64B64_SINT"`
	VkFormatR64G64B64Sfloat                      VkFormatR64G64B64Sfloat                      `json:"VK_FORMAT_R64G64B64_SFLOAT"`
	VkFormatR64G64B64A64Uint                     VkFormatR64G64B64A64Uint                     `json:"VK_FORMAT_R64G64B64A64_UINT"`
	VkFormatR64G64B64A64Sint                     VkFormatR64G64B64A64Sint                     `json:"VK_FORMAT_R64G64B64A64_SINT"`
	VkFormatR64G64B64A64Sfloat                   VkFormatR64G64B64A64Sfloat                   `json:"VK_FORMAT_R64G64B64A64_SFLOAT"`
	VkFormatB10G11R11UfloatPack32                VkFormatB10G11R11UfloatPack32                `json:"VK_FORMAT_B10G11R11_UFLOAT_PACK32"`
	VkFormatE5B9G9R9UfloatPack32                 VkFormatE5B9G9R9UfloatPack32                 `json:"VK_FORMAT_E5B9G9R9_UFLOAT_PACK32"`
	VkFormatD16Unorm                             VkFormatD16Unorm                             `json:"VK_FORMAT_D16_UNORM"`
	VkFormatD32Sfloat                            VkFormatD32Sfloat                            `json:"VK_FORMAT_D32_SFLOAT"`
	VkFormatS8Uint                               VkFormatS8Uint                               `json:"VK_FORMAT_S8_UINT"`
	VkFormatD16UnormS8Uint                       VkFormatD16UnormS8Uint                       `json:"VK_FORMAT_D16_UNORM_S8_UINT"`
	VkFormatD32SfloatS8Uint                      VkFormatD32SfloatS8Uint                      `json:"VK_FORMAT_D32_SFLOAT_S8_UINT"`
	VkFormatBc1RgbUnormBlock                     VkFormatBc1RgbUnormBlock                     `json:"VK_FORMAT_BC1_RGB_UNORM_BLOCK"`
	VkFormatBc1RgbSrgbBlock                      VkFormatBc1RgbSrgbBlock                      `json:"VK_FORMAT_BC1_RGB_SRGB_BLOCK"`
	VkFormatBc1RgbaUnormBlock                    VkFormatBc1RgbaUnormBlock                    `json:"VK_FORMAT_BC1_RGBA_UNORM_BLOCK"`
	VkFormatBc1RgbaSrgbBlock                     VkFormatBc1RgbaSrgbBlock                     `json:"VK_FORMAT_BC1_RGBA_SRGB_BLOCK"`
	VkFormatBc2UnormBlock                        VkFormatBc2UnormBlock                        `json:"VK_FORMAT_BC2_UNORM_BLOCK"`
	VkFormatBc2SrgbBlock                         VkFormatBc2SrgbBlock                         `json:"VK_FORMAT_BC2_SRGB_BLOCK"`
	VkFormatBc3UnormBlock                        VkFormatBc3UnormBlock                        `json:"VK_FORMAT_BC3_UNORM_BLOCK"`
	VkFormatBc3SrgbBlock                         VkFormatBc3SrgbBlock                         `json:"VK_FORMAT_BC3_SRGB_BLOCK"`
	VkFormatBc4UnormBlock                        VkFormatBc4UnormBlock                        `json:"VK_FORMAT_BC4_UNORM_BLOCK"`
	VkFormatBc4SnormBlock                        VkFormatBc4SnormBlock                        `json:"VK_FORMAT_BC4_SNORM_BLOCK"`
	VkFormatBc5UnormBlock                        VkFormatBc5UnormBlock                        `json:"VK_FORMAT_BC5_UNORM_BLOCK"`
	VkFormatBc5SnormBlock                        VkFormatBc5SnormBlock                        `json:"VK_FORMAT_BC5_SNORM_BLOCK"`
	VkFormatBc6HUfloatBlock                      VkFormatBc6HUfloatBlock                      `json:"VK_FORMAT_BC6H_UFLOAT_BLOCK"`
	VkFormatBc6HSfloatBlock                      VkFormatBc6HSfloatBlock                      `json:"VK_FORMAT_BC6H_SFLOAT_BLOCK"`
	VkFormatBc7UnormBlock                        VkFormatBc7UnormBlock                        `json:"VK_FORMAT_BC7_UNORM_BLOCK"`
	VkFormatBc7SrgbBlock                         VkFormatBc7SrgbBlock                         `json:"VK_FORMAT_BC7_SRGB_BLOCK"`
	VkFormatG8B8R83Plane420Unorm                 VkFormatG8B8R83Plane420Unorm                 `json:"VK_FORMAT_G8_B8_R8_3PLANE_420_UNORM"`
	VkFormatG8B8R82Plane420Unorm                 VkFormatG8B8R82Plane420Unorm                 `json:"VK_FORMAT_G8_B8R8_2PLANE_420_UNORM"`
	VkFormatG8B8R83Plane422Unorm                 VkFormatG8B8R83Plane422Unorm                 `json:"VK_FORMAT_G8_B8_R8_3PLANE_422_UNORM"`
	VkFormatG8B8R82Plane422Unorm                 VkFormatG8B8R82Plane422Unorm                 `json:"VK_FORMAT_G8_B8R8_2PLANE_422_UNORM"`
	VkFormatG8B8R83Plane444Unorm                 VkFormatG8B8R83Plane444Unorm                 `json:"VK_FORMAT_G8_B8_R8_3PLANE_444_UNORM"`
	VkFormatR10X6UnormPack16                     VkFormatR10X6UnormPack16                     `json:"VK_FORMAT_R10X6_UNORM_PACK16"`
	VkFormatR10X6G10X6Unorm2Pack16               VkFormatR10X6G10X6Unorm2Pack16               `json:"VK_FORMAT_R10X6G10X6_UNORM_2PACK16"`
	VkFormatG10X6B10X6R10X62Plane420Unorm3Pack16 VkFormatG10X6B10X6R10X62Plane420Unorm3Pack16 `json:"VK_FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16"`
	VkFormatR12X4UnormPack16                     VkFormatR12X4UnormPack16                     `json:"VK_FORMAT_R12X4_UNORM_PACK16"`
	VkFormatR12X4G12X4Unorm2Pack16               VkFormatR12X4G12X4Unorm2Pack16               `json:"VK_FORMAT_R12X4G12X4_UNORM_2PACK16"`
	VkFormatG12X4B12X4R12X42Plane420Unorm3Pack16 VkFormatG12X4B12X4R12X42Plane420Unorm3Pack16 `json:"VK_FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16"`
	VkFormatG16B16R163Plane420Unorm              VkFormatG16B16R163Plane420Unorm              `json:"VK_FORMAT_G16_B16_R16_3PLANE_420_UNORM"`
	VkFormatG16B16R162Plane420Unorm              VkFormatG16B16R162Plane420Unorm              `json:"VK_FORMAT_G16_B16R16_2PLANE_420_UNORM"`
	VkFormatG16B16R163Plane422Unorm              VkFormatG16B16R163Plane422Unorm              `json:"VK_FORMAT_G16_B16_R16_3PLANE_422_UNORM"`
	VkFormatG16B16R162Plane422Unorm              VkFormatG16B16R162Plane422Unorm              `json:"VK_FORMAT_G16_B16R16_2PLANE_422_UNORM"`
	VkFormatG16B16R163Plane444Unorm              VkFormatG16B16R163Plane444Unorm              `json:"VK_FORMAT_G16_B16_R16_3PLANE_444_UNORM"`
	VkFormatA4R4G4B4UnormPack16                  VkFormatA4R4G4B4UnormPack16                  `json:"VK_FORMAT_A4R4G4B4_UNORM_PACK16"`
	VkFormatA4B4G4R4UnormPack16                  VkFormatA4B4G4R4UnormPack16                  `json:"VK_FORMAT_A4B4G4R4_UNORM_PACK16"`
	VkFormatA1B5G5R5UnormPack16                  VkFormatA1B5G5R5UnormPack16                  `json:"VK_FORMAT_A1B5G5R5_UNORM_PACK16"`
	VkFormatA8Unorm                              VkFormatA8Unorm                              `json:"VK_FORMAT_A8_UNORM"`
}

type MinImageTransferGranularity struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Depth  int `json:"depth"`
}

type VkQueueFamilyProperties struct {
	MinImageTransferGranularity MinImageTransferGranularity `json:"minImageTransferGranularity"`
	QueueCount                  int                         `json:"queueCount"`
	QueueFlags                  []string                    `json:"queueFlags"`
	TimestampValidBits          int                         `json:"timestampValidBits"`
}

type VkQueueFamilyQueryResultStatusPropertiesKHR struct {
	QueryResultStatusSupport bool `json:"queryResultStatusSupport"`
}

type VkQueueFamilyVideoPropertiesKHR struct {
	VideoCodecOperations []any `json:"videoCodecOperations"`
}

type QueueFamiliesProperties struct {
	VkQueueFamilyProperties                     VkQueueFamilyProperties                     `json:"VkQueueFamilyProperties"`
	VkQueueFamilyQueryResultStatusPropertiesKHR VkQueueFamilyQueryResultStatusPropertiesKHR `json:"VkQueueFamilyQueryResultStatusPropertiesKHR"`
	VkQueueFamilyVideoPropertiesKHR             VkQueueFamilyVideoPropertiesKHR             `json:"VkQueueFamilyVideoPropertiesKHR"`
}

type VkVideoProfileInfoKHR struct {
	VideoCodecOperation string   `json:"videoCodecOperation"`
	ChromaSubsampling   []string `json:"chromaSubsampling"`
	LumaBitDepth        []string `json:"lumaBitDepth"`
	ChromaBitDepth      []string `json:"chromaBitDepth"`
}

type VkVideoDecodeAV1ProfileInfoKHR struct {
	StdProfile       string `json:"stdProfile"`
	FilmGrainSupport bool   `json:"filmGrainSupport"`
}

type Profile struct {
	VkVideoProfileInfoKHR          VkVideoProfileInfoKHR          `json:"VkVideoProfileInfoKHR"`
	VkVideoDecodeAV1ProfileInfoKHR VkVideoDecodeAV1ProfileInfoKHR `json:"VkVideoDecodeAV1ProfileInfoKHR"`
}

type PictureAccessGranularity struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type MinCodedExtent struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type MaxCodedExtent struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type StdHeaderVersion struct {
	ExtensionName string `json:"extensionName"`
	SpecVersion   int    `json:"specVersion"`
}

type VkVideoCapabilitiesKHR struct {
	Flags                             []string                 `json:"flags"`
	MinBitstreamBufferOffsetAlignment int                      `json:"minBitstreamBufferOffsetAlignment"`
	MinBitstreamBufferSizeAlignment   int                      `json:"minBitstreamBufferSizeAlignment"`
	PictureAccessGranularity          PictureAccessGranularity `json:"pictureAccessGranularity"`
	MinCodedExtent                    MinCodedExtent           `json:"minCodedExtent"`
	MaxCodedExtent                    MaxCodedExtent           `json:"maxCodedExtent"`
	MaxDpbSlots                       int                      `json:"maxDpbSlots"`
	MaxActiveReferencePictures        int                      `json:"maxActiveReferencePictures"`
	StdHeaderVersion                  StdHeaderVersion         `json:"stdHeaderVersion"`
}

type VkVideoDecodeCapabilitiesKHR struct {
	Flags []string `json:"flags"`
}

type VkVideoDecodeAV1CapabilitiesKHR struct {
	MaxLevel string `json:"maxLevel"`
}

type Capabilities struct {
	VkVideoCapabilitiesKHR          VkVideoCapabilitiesKHR          `json:"VkVideoCapabilitiesKHR"`
	VkVideoDecodeCapabilitiesKHR    VkVideoDecodeCapabilitiesKHR    `json:"VkVideoDecodeCapabilitiesKHR"`
	VkVideoDecodeAV1CapabilitiesKHR VkVideoDecodeAV1CapabilitiesKHR `json:"VkVideoDecodeAV1CapabilitiesKHR"`
}

type ComponentMapping struct {
	R string `json:"r"`
	G string `json:"g"`
	B string `json:"b"`
	A string `json:"a"`
}

type VkVideoFormatPropertiesKHR struct {
	Format           string           `json:"format"`
	ComponentMapping ComponentMapping `json:"componentMapping"`
	ImageCreateFlags []string         `json:"imageCreateFlags"`
	ImageType        string           `json:"imageType"`
	ImageTiling      string           `json:"imageTiling"`
	ImageUsageFlags  []string         `json:"imageUsageFlags"`
}

type VideoFormats struct {
	VkVideoFormatPropertiesKHR VkVideoFormatPropertiesKHR `json:"VkVideoFormatPropertiesKHR"`
}

type VideoProfiles struct {
	Profile      Profile        `json:"profile"`
	Capabilities Capabilities   `json:"capabilities"`
	Formats      []VideoFormats `json:"formats"`
}

type Device struct {
	Extensions              Extensions                `json:"extensions"`
	Features                Features                  `json:"features"`
	Properties              Properties                `json:"properties"`
	Formats                 Formats                   `json:"formats"`
	QueueFamiliesProperties []QueueFamiliesProperties `json:"queueFamiliesProperties"`
	VideoProfiles           []VideoProfiles           `json:"videoProfiles"`
}

type DeviceCapabilities struct {
	Device Device `json:"device"`
}

type History struct {
	Revision int    `json:"revision"`
	Date     string `json:"date"`
	Author   string `json:"author"`
	Comment  string `json:"comment"`
}

type VulkanDeviceDescription struct {
	Version      int       `json:"version"`
	APIVersion   string    `json:"api-version"`
	Label        string    `json:"label"`
	Description  string    `json:"description"`
	History      []History `json:"history"`
	Capabilities []string  `json:"capabilities"`
}

type Profiles map[string]VulkanDeviceDescription

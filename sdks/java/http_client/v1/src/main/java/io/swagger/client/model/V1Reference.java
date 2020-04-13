// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * OpenAPI spec version: 1.0.76
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


package io.swagger.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import io.swagger.client.model.V1DagReference;
import io.swagger.client.model.V1HubReference;
import io.swagger.client.model.V1PathReference;
import io.swagger.client.model.V1UrlReference;
import java.io.IOException;

/**
 * V1Reference
 */

public class V1Reference {
  @SerializedName("hub_reference")
  private V1HubReference hubReference = null;

  @SerializedName("dag_reference")
  private V1DagReference dagReference = null;

  @SerializedName("url_reference")
  private V1UrlReference urlReference = null;

  @SerializedName("path_reference")
  private V1PathReference pathReference = null;

  public V1Reference hubReference(V1HubReference hubReference) {
    this.hubReference = hubReference;
    return this;
  }

   /**
   * Get hubReference
   * @return hubReference
  **/
  @ApiModelProperty(value = "")
  public V1HubReference getHubReference() {
    return hubReference;
  }

  public void setHubReference(V1HubReference hubReference) {
    this.hubReference = hubReference;
  }

  public V1Reference dagReference(V1DagReference dagReference) {
    this.dagReference = dagReference;
    return this;
  }

   /**
   * Get dagReference
   * @return dagReference
  **/
  @ApiModelProperty(value = "")
  public V1DagReference getDagReference() {
    return dagReference;
  }

  public void setDagReference(V1DagReference dagReference) {
    this.dagReference = dagReference;
  }

  public V1Reference urlReference(V1UrlReference urlReference) {
    this.urlReference = urlReference;
    return this;
  }

   /**
   * Get urlReference
   * @return urlReference
  **/
  @ApiModelProperty(value = "")
  public V1UrlReference getUrlReference() {
    return urlReference;
  }

  public void setUrlReference(V1UrlReference urlReference) {
    this.urlReference = urlReference;
  }

  public V1Reference pathReference(V1PathReference pathReference) {
    this.pathReference = pathReference;
    return this;
  }

   /**
   * Get pathReference
   * @return pathReference
  **/
  @ApiModelProperty(value = "")
  public V1PathReference getPathReference() {
    return pathReference;
  }

  public void setPathReference(V1PathReference pathReference) {
    this.pathReference = pathReference;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Reference v1Reference = (V1Reference) o;
    return Objects.equals(this.hubReference, v1Reference.hubReference) &&
        Objects.equals(this.dagReference, v1Reference.dagReference) &&
        Objects.equals(this.urlReference, v1Reference.urlReference) &&
        Objects.equals(this.pathReference, v1Reference.pathReference);
  }

  @Override
  public int hashCode() {
    return Objects.hash(hubReference, dagReference, urlReference, pathReference);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Reference {\n");
    
    sb.append("    hubReference: ").append(toIndentedString(hubReference)).append("\n");
    sb.append("    dagReference: ").append(toIndentedString(dagReference)).append("\n");
    sb.append("    urlReference: ").append(toIndentedString(urlReference)).append("\n");
    sb.append("    pathReference: ").append(toIndentedString(pathReference)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}
